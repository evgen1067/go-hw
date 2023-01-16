Feature: list of events
  In order to be happy
  As a hungry gopher
  I need to be able to get an event list
  Scenario: not successful list (invalid event period)
    When I send "GET" request to "http://localhost:8888/events/list/fail"
    Then The response code should be 400
  Scenario: successful list
    When I send "GET" request to "http://localhost:8888/events/list/day?date=2025-01-16T19:00:00"
    Then The response code should be 200