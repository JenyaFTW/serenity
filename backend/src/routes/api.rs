use actix_web::web::{scope, get, post, ServiceConfig};

use crate::controllers::api::auth_controller::AuthController;

pub fn register(config: &mut ServiceConfig) {
    config.service(
        scope("/auth")
            .route("/me", get().to(AuthController::get_me))
            .route("/login", post().to(AuthController::post_login))
            .route("/signup", post().to(AuthController::post_signup)),
    );
}