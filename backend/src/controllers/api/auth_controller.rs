use actix_web::{Responder};

pub struct AuthController;

impl AuthController {
    pub async fn get_me() -> impl Responder {
        return "Hey";
    }

    pub async fn post_login() -> impl Responder {
        return "Hey";
    }

    pub async fn post_signup() -> impl Responder {
        return "Hey";
    }
}