export GO111MODULE := on

UID := demo
PORT := 1991
HOST := localhost
TOKEN_FILE := .idToken

NOTE_ID:=2
NOTE_TITLE:=title
NOTE_DESCRIPTION:=description

REF_ID:=1
REF_TITLE:=title
REF_DESCRIPTION:=description
REF_LINK:=https://example.com/

TAG_ID:=1
TAG_NAME:=tag

create-token:
	go run ./cmd/customtoken/main.go $(UID) $(TOKEN_FILE)

req-notes:
	curl -v $(HOST):$(PORT)/notes

req-img-pei:
	curl -v $(HOST):$(PORT)/img/pei.png

req-notes-get:
	curl -v $(HOST):$(PORT)/articles/$(NOTE_ID)

req-notes-post:
	curl -v -XPOST -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/notes -d '{"title": "$(NOTE_TITLE)", "description": "$(NOTE_DESCRIPTION)"}'

req-notes-update:
	curl -v -XPUT -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/notes/$(NOTE_ID) -d '{"title": "$(ARTICLE_TITLE)", "body": "$(ARTICLE_BODY)"}'

req-notes-delete:
	curl -v -XDELETE -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/notes/$(NOTE_ID)


req-refs-get:
	curl -v $(HOST):$(PORT)/refs/$(REF_ID)

req-refs-post:
	curl -v -XPOST -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/notes/$(NOTE_ID)/refs -d '{"title": "$(REF_TITLE)", "description": "$(REF_DESCRIPTION)", "link": "${REF_LINK}"}'

req-refs-update:
	curl -v -XPUT -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/refs/$(REF_ID) -d '{"title": "$(REF_TITLE)", "description": "$(REF_DESCRIPTION)"}'

req-refs-delete:
	curl -v -XDELETE -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/refs/$(REF_ID)


req-tags:
	curl -v $(HOST):$(PORT)/tags

req-tags-get:
	curl -v $(HOST):$(PORT)/tags/$(TAG_ID)

req-tags-post:
	curl -v -XPOST -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/tags -d '{"name": "$(TAG_NAME)"}'

req-tags-update:
	curl -v -XPUT -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/tags/$(TAG_ID) -d '{"name": "$(TAG_NAME)"}'

req-tags-delete:
	curl -v -XDELETE -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/tags/$(TAG_ID)


req-note-tag-post:
	curl -v -XPOST -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/notes/${NOTE_ID}/tags/${TAG_ID}

req-note-tag-delete:
	curl -v -XDELETE -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/notes/${NOTE_ID}/tags/${TAG_ID}


req-public:
	curl -v $(HOST):$(PORT)/public

req-private:
	curl -v -H "Authorization: Bearer $(shell cat ./$(TOKEN_FILE))" $(HOST):$(PORT)/private

database-init:
	make -C ../database init
