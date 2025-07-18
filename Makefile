INSTALL := .repocli/install

.PHONY: dev
dev:
	@clear
	@docker compose down arch
	@docker compose up arch -d
	@docker compose exec arch bash -c 'cat $$HOME/$(INSTALL) | bash'


.PHONY: prod
prod:
	@clear
	@docker compose down arch
	@docker compose up arch -d
	@docker compose exec arch bash -c 'wget -qO- "https://raw.githubusercontent.com/jgttech/repo/refs/heads/main/install" | bash'

.PHONY: ssh
ssh:
	@docker compose exec arch bash
