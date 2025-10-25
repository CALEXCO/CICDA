-- PostgreSQL init script for the 'cicda' database
-- NOTE: connect with psql -U cicda -d cicda -f init.sql (do not include a 'USE' statement)

-- Create table 'project'
CREATE TABLE IF NOT EXISTS project (
    project_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    leader TEXT,
    description TEXT,
    date_of_creation DATE NOT NULL DEFAULT CURRENT_DATE
);

-- Insert sample rows: columns are (name, leader, description)
INSERT INTO project (name, leader, description) VALUES
('Project Aurora', 'María Gómez', 'Real-time monitoring platform'),
('Project Boreal', 'Juan Pérez', 'CI/CD deployment automation'),
('Project Cobalt', 'Ana Ruiz', 'User and permissions management API'),
('Project Delta', 'Luis Martínez', 'Reporting and analytics system'),
('Project Epsilon', 'Sofía López', 'Mobile app for task tracking');
