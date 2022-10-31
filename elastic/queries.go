package elastic

import "fmt"

func MessageQuery(msgId string) string {
	query := fmt.Sprintf(`{
		"query": {
			"term": {
				"id": "%s"
			}
		}
	}`, msgId)

	return query
}


func UserQuery(userId string) string {
	query := fmt.Sprintf(`{
		"query": {
			"term": {
				"id": "%s"
			}
		}
	}`, userId)

	return query
}