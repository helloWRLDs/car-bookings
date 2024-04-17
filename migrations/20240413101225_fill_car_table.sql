-- +goose Up
-- +goose StatementBegin
INSERT INTO cars (vendor, model, year, is_available, body, engine_capacity, mileage, color) VALUES
('Toyota', 'Corolla', 2018, true, 'Sedan', 1.8, 50000, 'Red'),
('Honda', 'Civic', 2020, true, 'Sedan', 1.5, 40000, 'Blue'),
('Ford', 'F-150', 2019, false, 'Truck', 5.0, 60000, 'Black'),
('Chevrolet', 'Camaro', 2021, true, 'Coupe', 6.2, 30000, 'Yellow'),
('BMW', 'X5', 2017, false, 'SUV', 3.0, 80000, 'White'),
('Audi', 'A4', 2019, true, 'Sedan', 2.0, 45000, 'Silver'),
('Mercedes-Benz', 'E-Class', 2016, false, 'Sedan', 3.0, 70000, 'Grey'),
('Jeep', 'Wrangler', 2020, true, 'SUV', 3.6, 35000, 'Green'),
('Volkswagen', 'Golf', 2019, true, 'Hatchback', 1.8, 40000, 'Blue'),
('Nissan', 'Altima', 2020, false, 'Sedan', 2.5, 55000, 'Black'),
('Hyundai', 'Tucson', 2018, true, 'SUV', 2.4, 48000, 'Red'),
('Kia', 'Sorento', 2019, true, 'SUV', 3.3, 42000, 'Silver'),
('Subaru', 'Outback', 2021, true, 'Wagon', 2.5, 30000, 'White'),
('Mazda', 'CX-5', 2020, false, 'SUV', 2.5, 40000, 'Grey'),
('Volvo', 'XC90', 2017, true, 'SUV', 2.0, 60000, 'Black'),
('Ford', 'Focus', 2017, true, 'Hatchback', 2.0, 70000, 'Red'),
('Chevrolet', 'Equinox', 2019, false, 'SUV', 2.0, 55000, 'Blue'),
('Toyota', 'Rav4', 2018, true, 'SUV', 2.5, 45000, 'Silver'),
('Honda', 'Accord', 2019, true, 'Sedan', 1.5, 50000, 'White'),
('Jeep', 'Grand Cherokee', 2016, false, 'SUV', 3.6, 80000, 'Black'),
('Subaru', 'Forester', 2019, true, 'SUV', 2.5, 40000, 'Grey'),
('Nissan', 'Rogue', 2020, true, 'SUV', 2.5, 35000, 'Red'),
('Hyundai', 'Elantra', 2021, true, 'Sedan', 2.0, 30000, 'Blue'),
('Kia', 'Optima', 2018, false, 'Sedan', 2.4, 60000, 'Silver'),
('Mazda', '3', 2020, true, 'Hatchback', 2.5, 35000, 'Red');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM cars;
-- +goose StatementEnd
