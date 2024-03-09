# Go_midterm

# Historical Figures REST API
```
GET /v1/healthcheck
POST /v1/figures
GET /v1/figures/:id
PUT /v1/figures/:id
DELETE /v1/figures/:id
```
# DB Structure
```
Table figures {
    id bigserial [primary key]
    created_at timestamp
    name text
    years_of_life text
    description text
    version integer
}
```
