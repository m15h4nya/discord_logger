package elastic

const ignoreError = "resource_already_exists_exception"

const messagesIndex = `{
	"mappings": {
		"properties": {
			"id": {
				"type": "keyword"
			},
			"channel_id": {
				"type": "keyword"
			},
			"guild_id": {
				"type": "keyword"
			},
			"content": {
				"type": "text"
			},
			"attachment": {
				"file_name": {
					"type": "keyword"
				},
				"bytes": {
					"type": "keyword"
				},
				"file_extension": {
					"type": "keyword"
				}
			},
			"author": {
				"type": "keyword"
			},
			"reference_id": {
				"type": "keyword"
			}
		}
	},
	"settings": {
		"number_of_replicas": "0"
		}
	}
}`

const userIndex = `{
	"mappings": {
		"properties": {
			"id": {
				"type": "keyword"
			},
			"username": {
				"type": "keyword"
			},
			"avatar": {
				"bytes": {
					"type": "keyword"
				},
				"file_extension": {
					"type": "keyword"
				}
			},
			"author": {
				"type": "keyword"
			}
		}
	},
	"settings": {
		"number_of_replicas": "0"
		}
	}
}`
