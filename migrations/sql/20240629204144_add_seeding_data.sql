-- +goose Up
-- +goose StatementBegin
-- Insert data into cart_items
INSERT INTO cart_items (user_id, product_id, qty)
VALUES
    ('a4b7a3f1-751a-4a10-b506-99202581427b', '550e8400-e29b-41d4-a716-446655440001', 2),
    ('a4b7a3f1-751a-4a10-b506-99202581427b', '550e8400-e29b-41d4-a716-446655440002', 1),
    ('a4b7a3f1-751a-4a10-b506-99202581427b', '550e8400-e29b-41d4-a716-446655440003', 3),
    ('c396f23e-a097-476d-aae5-cfc9973634f3', '550e8400-e29b-41d4-a716-446655440005', 4),
    ('c396f23e-a097-476d-aae5-cfc9973634f3', '550e8400-e29b-41d4-a716-446655440006', 5);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
