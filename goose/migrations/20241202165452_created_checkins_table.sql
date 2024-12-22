-- +goose Up
-- +goose StatementBegin

-- CreateTable
CREATE TABLE "checkins" (
    "id_checkin" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "validated_at" TIMESTAMP(3),
    "id_user" TEXT NOT NULL,
    "id_gym" TEXT NOT NULL,

    CONSTRAINT "checkins_pkey" PRIMARY KEY ("id_checkin")
);

-- AddForeignKey
ALTER TABLE "checkins" ADD CONSTRAINT "checkins_id_user_fkey" FOREIGN KEY ("id_user") REFERENCES "users"("id_user") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "checkins" ADD CONSTRAINT "checkins_id_gym_fkey" FOREIGN KEY ("id_gym") REFERENCES "gyms"("id_gym") ON DELETE RESTRICT ON UPDATE CASCADE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "checkins"
-- +goose StatementEnd
