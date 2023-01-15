package grpcapi

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"github.com/evgen1067/hw12_13_14_15_calendar/api"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/config"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository"
	"github.com/evgen1067/hw12_13_14_15_calendar/internal/repository/memory"
	data "github.com/evgen1067/hw12_13_14_15_calendar/internal/server/grpcapi/transformer"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestHandlers(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client, closer := server(ctx)
	defer closer()

	dateStart := time.Now()

	event := data.TransformEventToPb(repository.Event{
		ID:          0,
		Title:       "Title",
		Description: "Description",
		DateStart:   dateStart,
		DateEnd:     dateStart.AddDate(0, 0, 1),
		NotifyIn:    0,
		OwnerID:     0,
	})
	createRequest := &api.CreateRequest{Event: event}
	createRequests := make([]*api.CreateRequest, 0)
	for i := 1; i < 10; i++ {
		e := data.TransformEventToPb(repository.Event{
			ID:          repository.EventID(i),
			Title:       "Title",
			Description: "Description",
			DateStart:   dateStart.AddDate(0, 0, i*1),
			DateEnd:     dateStart.AddDate(0, 0, i*2),
			NotifyIn:    0,
			OwnerID:     0,
		})

		createRequests = append(createRequests, &api.CreateRequest{Event: e})
	}
	updateRequest := &api.UpdateRequest{
		Id:    0,
		Event: event,
	}
	deleteRequest := &api.DeleteRequest{Id: 0}

	listRequest := &api.ListRequest{
		Date: &timestamp.Timestamp{Seconds: dateStart.Unix(), Nanos: int32(dateStart.Nanosecond())},
	}

	t.Run("Test GRPC create, update, delete", func(t *testing.T) {
		create, err := client.Create(ctx, createRequest)
		require.NoError(t, err)
		require.Equal(t, 0, int(create.Id))

		_, err = client.Create(ctx, createRequest)
		require.Error(t, err)

		update, err := client.Update(ctx, updateRequest)
		require.NoError(t, err)
		require.Equal(t, 0, int(update.Id))

		updateRequest.Id = 1
		_, err = client.Update(ctx, updateRequest)
		require.Error(t, err)

		remove, err := client.Delete(ctx, deleteRequest)
		require.NoError(t, err)
		require.Equal(t, 0, int(remove.Id))

		_, err = client.Delete(ctx, deleteRequest)
		require.Error(t, err)
	})
	t.Run("Test GRPC list", func(t *testing.T) {
		for i, v := range createRequests {
			create, err := client.Create(ctx, v)
			require.NoError(t, err)
			require.Equal(t, i+1, int(create.Id))
		}

		list, err := client.DayList(ctx, listRequest)
		require.NoError(t, err)
		require.Equal(t, 0, len(list.Event))

		list, err = client.WeekList(ctx, listRequest)
		require.NoError(t, err)
		require.Equal(t, 6, len(list.Event))

		list, err = client.MonthList(ctx, listRequest)
		require.NoError(t, err)
		require.Equal(t, 9, len(list.Event))
	})
}

func server(ctx context.Context) (api.EventServiceClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	cfg, _ := config.InitConfig("../../../configs/local.json")
	repo := memory.NewRepo()
	grpcSrv := InitGRPC(ctx, repo, cfg)

	baseServer := grpc.NewServer()

	api.RegisterEventServiceServer(baseServer, grpcSrv)
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := api.NewEventServiceClient(conn)

	return client, closer
}