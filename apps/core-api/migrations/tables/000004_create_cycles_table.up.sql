CREATE TABLE IF NOT EXISTS cycles (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    sub_place_id INT NOT NULL,
    commodity_type VARCHAR(50) NOT NULL,
    CHECK (
        commodity_type IN ('fish', 'plant')
    ),
    commodity_name VARCHAR(255) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE, -- NULLABLE
    status VARCHAR(20) DEFAULT 'ongoing',
    CHECK (
        status IN (
            'ongoing',
            'finished',
            'failed'
        )
    ),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_sub_place FOREIGN KEY (sub_place_id) REFERENCES sub_places (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_cycles_sub_place_id ON cycles (sub_place_id);

CREATE INDEX IF NOT EXISTS idx_cycles_status ON cycles (status);

CREATE INDEX IF NOT EXISTS idx_cycles_start_date ON cycles (start_date);