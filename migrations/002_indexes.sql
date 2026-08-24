CREATE INDEX IF NOT EXISTS idx_versions_tenant_status ON agent_versions(tenant_id, status, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_grants_expiry ON tool_grants(tenant_id, status, expires_at);
CREATE INDEX IF NOT EXISTS idx_steps_run_status ON run_steps(run_id, status, sequence_no);
CREATE INDEX IF NOT EXISTS idx_schedules_due ON schedules(status, next_run_at);
CREATE INDEX IF NOT EXISTS idx_usage_tenant_time ON usage_ledger(tenant_id, created_at DESC);

