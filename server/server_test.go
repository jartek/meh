package server_test

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	. "github.com/jartek/worldcup/server"
	"github.com/martini-contrib/render"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func sliceFromJSON(data []byte) []interface{} {
	var result interface{}
	json.Unmarshal(data, &result)
	return result.([]interface{})
}

func mapFromJSON(data []byte) map[string]interface{} {
	var result interface{}
	json.Unmarshal(data, &result)
	return result.(map[string]interface{})
}

var _ = Describe("Server", func() {
	var db *sql.DB
	var server Server
	var request *http.Request
	var response *httptest.ResponseRecorder

	BeforeEach(func() {
		os.Setenv("DB_NAME", "world_cup_test")
		os.Setenv("DB_PASSWORD", "test")
		os.Setenv("DB_USER", "jartek")
		os.Setenv("DB_SSLMODE", "disable")
		db = SetupDB()
		server = NewServer(db)
		server.Use(render.Renderer())
		response = httptest.NewRecorder()
	})

	AfterEach(func() {
		db.Query("DROP DATABASE world_cup_test")
	})

	Describe("GET /teams", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/teams", nil)
		})
		It("Should return a status code of 200", func() {
			server.ServeHTTP(response, request)
			Expect(response.Code).To(Equal(200))
		})
		It("Should return a list of all teams", func() {
			var count int
			server.ServeHTTP(response, request)
			teams := sliceFromJSON(response.Body.Bytes())
			err := db.QueryRow("SELECT COUNT(*) FROM teams").Scan(&count)
			if err != nil {
				log.Fatalln(err)
			}
			Expect(len(teams)).To(Equal(count))
		})
	})

	Describe("GET /teams/:id", func() {
		Context("Valid ID", func() {
			BeforeEach(func() {
				request, _ = http.NewRequest("GET", "/teams/1", nil)
			})

			It("Should return a status code of 200", func() {
				server.ServeHTTP(response, request)
				Expect(response.Code).To(Equal(200))
			})

			It("Should return brazil", func() {
				server.ServeHTTP(response, request)
				team := mapFromJSON(response.Body.Bytes())
				Expect(team["Id"]).To(Equal(float64(1)))
				Expect(team["Name"]).To(Equal("Brazil"))
				Expect(team["NickName"]).To(Equal("Canarinho"))
			})

		})
		Context("Invalid ID", func() {
			BeforeEach(func() {
				request, _ = http.NewRequest("GET", "/teams/11231231", nil)
			})

			It("Should return a status code of 404", func() {
				server.ServeHTTP(response, request)
				Expect(response.Code).To(Equal(404))
			})
		})
	})

	Describe("GET /stadiums", func() {
		BeforeEach(func() {
			request, _ = http.NewRequest("GET", "/stadiums", nil)
		})

		It("Should return a status code of 200", func() {
			server.ServeHTTP(response, request)
			Expect(response.Code).To(Equal(200))
		})

		It("Should return a list of stadiums", func() {
			var count int
			server.ServeHTTP(response, request)
			stadiums := sliceFromJSON(response.Body.Bytes())
			err := db.QueryRow("SELECT COUNT(*) FROM stadiums").Scan(&count)
			if err != nil {
				log.Fatalln(err)
			}
			Expect(len(stadiums)).To(Equal(count))
		})
	})

	Describe("GET /stadiums/:id", func() {
		Context("Valid ID", func() {
			BeforeEach(func() {
				request, _ = http.NewRequest("GET", "/stadiums/1", nil)
			})

			It("Should return a status code of 200", func() {
				server.ServeHTTP(response, request)
				Expect(response.Code).To(Equal(200))
			})

			It("Should return the first record", func() {
				server.ServeHTTP(response, request)
				stadium := mapFromJSON(response.Body.Bytes())
				Expect(stadium["Id"]).To(Equal(float64(1)))
				Expect(stadium["Name"]).To(Equal("Arena De Sao Paulo"))
			})

		})
		Context("Invalid ID", func() {
			BeforeEach(func() {
				request, _ = http.NewRequest("GET", "/stadiums/11231231", nil)
			})

			It("Should return a status code of 404", func() {
				server.ServeHTTP(response, request)
				Expect(response.Code).To(Equal(404))
			})
		})
	})

})
