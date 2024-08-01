#[get("/new")]
pub fn new() -> String {
    format!("todo new host")
}

#[get("/all")]
pub fn all() -> String {
    format!("todo all hosts")
}

#[get("/edit/<id>")]
pub fn edit(id: i32) -> String {
    format!("todo edit host {}", id)
}

#[get("/del/<id>")]
pub fn del(id: i32) -> String {
    format!("todo del host {}", id)
}
