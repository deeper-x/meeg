#[macro_use]
extern crate rocket;

mod router {
    pub mod db;
    pub mod host;
    pub mod migration;
}

#[launch]
fn rocket() -> _ {
    rocket::build()
        .mount(
            "/host",
            routes![
                router::host::new,
                router::host::all,
                router::host::edit,
                router::host::del,
            ],
        )
        .mount(
            "/db",
            routes![
                router::db::new,
                router::db::all,
                router::db::edit,
                router::db::del,
            ],
        )
        .mount(
            "/migration",
            routes![
                router::migration::new,
                router::migration::all,
                router::migration::register,
                router::migration::del,
            ],
        )
}
