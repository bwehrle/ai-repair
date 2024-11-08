
.PHONY: healtcheck-openai-4o
healtcheck-openai-4o:
	@echo "Running $(@)"
	@aider --model openai/gpt-4o --dry-run --load scenarios/healthchecks/scenario.cmd

.PHONY: healtcheck-anthropic-haiku
healtcheck-anthropic-sonnet:
	@echo "Running $(@)"
	@aider --model anthropic/claude-3-5-sonnet-20241022 --dry-run --load scenarios/healthchecks/scenario.cmd

