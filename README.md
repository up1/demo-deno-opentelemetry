# Deno with OpenTelemetry
* https://docs.deno.com/runtime/fundamentals/open_telemetry/

## 1. Create LGTM stack
```
$docker compose up -d lgtm
```

## 2. Create API
```
$docker compose build api
$docker compose up -d api

$docker compose ps       
NAME              IMAGE                     COMMAND                  SERVICE   CREATED         STATUS         PORTS
demo-otel-api-1   demo-otel-api             "/tini -- docker-ent…"   api       4 seconds ago   Up 4 seconds   0.0.0.0:8000->8000/tcp
lgtm              grafana/otel-lgtm:0.8.1   "/bin/sh -c /otel-lg…"   lgtm      4 minutes ago   Up 4 minutes   0.0.0.0:3000->3000/tcp, 0.0.0.0:4317-4318->4317-4318/tcp
```

List of API
* http://localhost:8000/

Access to Grafana
* http://localhost:3000/
  * Explore


## 3. Working with Go-Auto instrumentation
```
$docker compose up -d lgtm
$docker compose ps

$docker compose up -d go-api --build
$docker compose ps

$docker compose up -d go-auto
$docker compose ps
```

