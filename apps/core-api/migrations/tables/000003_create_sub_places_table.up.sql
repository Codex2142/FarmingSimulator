CREATE TABLE IF NOT EXISTS sub_places (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    farm_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    CHECK (
        type IN (
            'pond',
            'plant_bed',
            'greenhouse'
        )
    ),
    size NUMERIC(10, 2),
    status VARCHAR(20) DEFAULT 'active',
    CHECK (
        status IN ('active', 'inactive')
    ),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_farm FOREIGN KEY (farm_id) REFERENCES farms (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_sub_places_farm_id ON sub_places (farm_id);

CREATE INDEX IF NOT EXISTS idx_sub_places_status ON sub_places (status);