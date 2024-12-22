-- +goose Up
-- +goose StatementBegin
-- CreateTable
CREATE TABLE "gyms" (
    "id_gym" TEXT NOT NULL,
    "gym_name" TEXT NOT NULL,
    "description" TEXT,
    "phone" TEXT,
    "latitude" DECIMAL(65,30) NOT NULL,
    "longitude" DECIMAL(65,30) NOT NULL,

    CONSTRAINT "gym_pkey" PRIMARY KEY ("id_gym")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "gyms"
-- +goose StatementEnd
