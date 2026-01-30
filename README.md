## Documentation

- [Design & Implementation Overview](docs/design.md)

## Build and Run
- To build the project
```
    ./scripts/build.sh
```
This will create a docker image **monopoly-go-bank-heist:latest**

To run API server from docker container
```
    ./scripts/run.sh
```

## API Endpoint Test

### Prediction Endpoint
* Request
```
curl -X POST http://localhost:8080/api/predict   -H "Content-Type: application/json"   -d '{
    "doors": [
      { "row": 1, "col": 1, "outcome": 1 },
      { "row": 2, "col": 2, "outcome": 3 }
    ]
  }'
```
* Response
```
    {"recommendation":{"row":1,"col":2,"outcome":1,"probability":0.2},"doors":[{"row":1,"col":2,"outcome":1,"probability":0.2},{"row":1,"col":3,"outcome":1,"probability":0.2},{"row":1,"col":4,"outcome":1,"probability":0.2},{"row":2,"col":1,"outcome":1,"probability":0.2},{"row":2,"col":3,"outcome":1,"probability":0.2},{"row":2,"col":4,"outcome":1,"probability":0.2},{"row":3,"col":1,"outcome":1,"probability":0.2},{"row":3,"col":2,"outcome":1,"probability":0.2},{"row":3,"col":3,"outcome":1,"probability":0.2},{"row":3,"col":4,"outcome":1,"probability":0.2}]}
```

### Update Endpoint
* Request
```
curl -X POST http://localhost:8080/api/update \
  -H "Content-Type: application/json" \
  -d '{
    "doors": [
      { "row": 1, "col": 1, "outcome": 2 },
      { "row": 1, "col": 2, "outcome": 2 },
      { "row": 1, "col": 3, "outcome": 1 },
      { "row": 1, "col": 4, "outcome": 1 },

      { "row": 2, "col": 1, "outcome": 2 },
      { "row": 2, "col": 2, "outcome": 3 },
      { "row": 2, "col": 3, "outcome": 3 },
      { "row": 2, "col": 4, "outcome": 2 },

      { "row": 3, "col": 1, "outcome": 1 },
      { "row": 3, "col": 2, "outcome": 3 },
      { "row": 3, "col": 3, "outcome": 1 },
      { "row": 3, "col": 4, "outcome": 2 }
    ]
  }'

```
* Response
```
    {"status":"pattern frequency updated"}
```
