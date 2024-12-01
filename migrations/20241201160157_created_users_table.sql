-- +goose Up
-- +goose StatementBegin
-- CreateEnum
CREATE TYPE "Role" AS ENUM ('ADMIN', 'MEMBER');

-- CreateTable
CREATE TABLE "users" (
    "id_user" TEXT NOT NULL,
    "user_name" TEXT NOT NULL,
  	"password_hash" TEXT NOT NULL,
  	"role" "Role" NOT NULL DEFAULT 'MEMBER',
    "email" TEXT NOT NULL,
  	"created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "users_pkey" PRIMARY KEY ("id_user")
);

-- CreateIndex
CREATE UNIQUE INDEX "users_email_key" ON "users"("email");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "users"
-- +goose StatementEnd
