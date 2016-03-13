# greenapple
Small server for keeping track of learnt algorithms.

### Endoints
* **[GET] /** - Returns all the algorithms stored in a json form.
* **[GET] /algo/{id}** 
* **[POST] /algo** - Creates a new algorithm definition in the database.
```bash
# e.g.
curl  \
  -H "Content-Type: application/json" \
  -d '{"name": "Dijkstra", "desc": "Single source, shortest path algorithm"}' 
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
