# Learn 5 Minutes
Small server for keeping track of learnt algorithms.

### Endoints
* **[GET] /** - Returns all the algorithms stored in a json form.
* **[GET] /algo/{id}** 
* **[POST] /algo** - Creates a new algorithm definition in the database.
```bash
# e.g.
curl  \
  -H "Content-Type: application/json" \
  -X POST \
  -d '{"name": "Dijkstra", "desc": "Single source, shortest path algorithm"}' \
  localhost:8080/algo
```
* **[PUT] /algo** - Updates an existing algorithm based on algorithm id.
```bash
# e.g.
# Do not forget to specify algorithm ID, otherwise it will not know what to update.
curl  \
  -H "Content-Type: application/json" \
  -X PUT \
  -d '{"id":1", name": "Dijkstra updated", "desc": "UPDATED DESCRIPTION"}' \
  localhost:8080/algo
```
* **[GET] /random** - Returns a random algorithm from the server.

---

### Algo Data Structure
```json
{
  "id": 7,
  "name": "Dijkstra",
  "desc": "Single source, shortest path algorithm on graphs"
}
```

### Run
```bash
go get
go build
./greenapple -port=8080 &
```
