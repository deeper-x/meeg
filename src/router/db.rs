#[get("/new/<id>")]
pub fn new(id: i32) -> String {
    format!("todo new db {}", id)
}

#[get("/all")]
pub fn all() -> String {
    format!("todo all")
}

#[get("/edit/<id>")]
pub fn edit(id: i32) -> String {
    format!("todo edit {}", id)
}

#[get("/del/<id>")]
pub fn del(id: i32) -> String {
    format!("todo del {}", id)
}
