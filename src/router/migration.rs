#[get("/new/<id>")]
pub fn new(id: i32) -> String {
    format!("todo new migration for db {}", id)
}

#[get("/all/<db_id>")]
pub fn all(db_id: i32) -> String {
    format!("todo all migrations for db {}", db_id)
}

#[get("/del/<id>")]
pub fn del(id: i32) -> String {
    format!("todo del {}", id)
}

#[get("/register/<id>/<host_id>")]
pub fn register(id: i32, host_id: i32) -> String {
    format!("todo register migration {} w/ host {}", id, host_id)
}
