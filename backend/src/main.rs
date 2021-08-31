use actix_web::{web, HttpServer, App};
use actix_files as fs;

use {routes::api::register as api_routes};

mod routes;
mod controllers;

async fn vue_index() -> actix_files::NamedFile{
    actix_files::NamedFile::open("../frontend/dist/index.html").unwrap()
}

#[actix_web::main]
async fn main() -> std::io::Result<()>
{
    HttpServer::new(|| {
        App::new()
            .service(
                web::scope("/api").configure(api_routes)
            )
            .service(fs::Files::new("/", "../frontend/dist").index_file("index.html"))
            .default_service(web::resource("").route(web::get().to(vue_index)))
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}