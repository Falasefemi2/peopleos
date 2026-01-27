-- Memo Types
CREATE TABLE IF NOT EXISTS memo_types (
    id SERIAL PRIMARY KEY,
    tenant_id INTEGER NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (tenant_id) REFERENCES tenants(id) ON DELETE CASCADE,
    UNIQUE(tenant_id, name)
);

-- Memos
CREATE TABLE IF NOT EXISTS memos (
    id SERIAL PRIMARY KEY,
    tenant_id INTEGER NOT NULL,
    employee_id INTEGER NOT NULL,
    memo_type_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) DEFAULT 'pending',
    approval_workflow_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (tenant_id) REFERENCES tenants(id) ON DELETE CASCADE,
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE,
    FOREIGN KEY (memo_type_id) REFERENCES memo_types(id) ON DELETE RESTRICT,
    FOREIGN KEY (approval_workflow_id) REFERENCES approval_workflows(id) ON DELETE SET NULL
);
