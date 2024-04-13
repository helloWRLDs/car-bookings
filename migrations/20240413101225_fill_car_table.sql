-- +goose Up
-- +goose StatementBegin
INSERT INTO cars (vendor, model, year, isAvailable, body) VALUES
('Toyota', 'Corolla', 2018, true, 'Sedan'),
('Honda', 'Civic', 2020, true, 'Sedan'),
('Ford', 'F-150', 2019, false, 'Truck'),
('Chevrolet', 'Camaro', 2021, true, 'Coupe'),
('BMW', 'X5', 2017, false, 'SUV'),
('Audi', 'A4', 2019, true, 'Sedan'),
('Mercedes-Benz', 'E-Class', 2016, false, 'Sedan'),
('Jeep', 'Wrangler', 2020, true, 'SUV'),
('Volkswagen', 'Golf', 2019, true, 'Hatchback'),
('Nissan', 'Altima', 2020, false, 'Sedan'),
('Hyundai', 'Tucson', 2018, true, 'SUV'),
('Kia', 'Sorento', 2019, true, 'SUV'),
('Subaru', 'Outback', 2021, true, 'Wagon'),
('Mazda', 'CX-5', 2020, false, 'SUV'),
('Volvo', 'XC90', 2017, true, 'SUV'),
('Ford', 'Focus', 2017, true, 'Hatchback'),
('Chevrolet', 'Equinox', 2019, false, 'SUV'),
('Toyota', 'Rav4', 2018, true, 'SUV'),
('Honda', 'Accord', 2019, true, 'Sedan'),
('Jeep', 'Grand Cherokee', 2016, false, 'SUV'),
('Subaru', 'Forester', 2019, true, 'SUV'),
('Nissan', 'Rogue', 2020, true, 'SUV'),
('Hyundai', 'Elantra', 2021, true, 'Sedan'),
('Kia', 'Optima', 2018, false, 'Sedan'),
('Mazda', '3', 2020, true, 'Hatchback');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM cars;
-- +goose StatementEnd
