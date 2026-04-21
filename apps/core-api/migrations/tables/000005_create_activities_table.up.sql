CREATE TABLE IF NOT EXISTS activities (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    cycle_id INT NOT NULL,
    type VARCHAR(50) NOT NULL,
    CHECK (
        type IN (
            'feeding',
            'fertilizing',
            'cleaning',
            'note'
        )
    ),
    description TEXT NOT NULL,
    quantity NUMERIC(10, 2),
    unit VARCHAR(50),
    created_by INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_cycle FOREIGN KEY (cycle_id) REFERENCES cycles (id) ON DELETE CASCADE,
    CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users (id) ON DELETE SET NULL
);

CREATE INDEX IF NOT EXISTS idx_activities_cycle_id ON activities (cycle_id);

CREATE INDEX IF NOT EXISTS idx_activities_created_by ON activities (created_by);

CREATE INDEX IF NOT EXISTS idx_activities_type ON activities(type);

CREATE INDEX IF NOT EXISTS idx_activities_created_at ON activities (created_at);