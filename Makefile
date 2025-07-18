.PHONY: arch
arch:
	@clear
	@docker compose down arch
	@docker compose up arch -d
	@docker compose exec arch bash -c 'wget -qO- "https://raw.githubusercontent.com/jgttech/repo/refs/heads/main/install" | bash'

.PHONY: bash
bash:
	@docker compose exec arch bash
