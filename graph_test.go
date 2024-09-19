package dndgen_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/suite"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/migrate"
	gen "github.com/ecshreve/dndgen/graph"
	"github.com/ecshreve/dndgen/internal/popper"
)

type EndToEndSuite struct {
	suite.Suite
	gql *client.Client
	ent *ent.Client
}

func (s *EndToEndSuite) SetupTest() {
	cl, err := ent.Open("sqlite3", "file:ent?cache=shared&mode=memory&_fk=1")
	s.Require().NoError(err)

	err = cl.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))
	s.Require().NoError(err)

	p := popper.NewPopper(context.Background(), cl)
	err = p.PopulateAll(context.Background())
	s.Require().NoError(err)

	srv := handler.NewDefaultServer(gen.NewSchema(cl))
	gc := client.New(srv)

	s.gql = gc
	s.ent = cl
}

func TestEndToEnd(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) TestQuery() {
	testCases := []struct {
		desc  string
		query string
	}{
		{
			desc: "abilityScores",
			query: `query {
				abilityScores {
					indx
					name
					fullName
					desc
				}
			}
			`,
		},
		{
			desc: "alignments",
			query: `query {
				alignments {
					indx
					name
					desc
					abbr
				}
			}
			`,
		},
		{
			desc: "languages",
			query: `query {
				languages {
					indx
					name
					desc
					languageType
					script
				}
			}
			`,
		},
		{
			desc: "races",
			query: `query {
				races {
					indx
					name
					speed
					size
					sizeDesc
					ageDesc
					alignmentDesc
					languageDesc
				}
			}
			`,
		},
		{
			desc: "skills",
			query: `query {
				skills {
					indx
					name
					desc
					abilityScore {
						indx
					}
				}
			}
			`,
		},
	}

	for _, tc := range testCases {
		var resp map[string]interface{}
		err := s.gql.Post(tc.query, &resp)
		s.Require().NoError(err)

		snaps.WithConfig(snaps.Filename(fmt.Sprintf("%s", tc.desc))).MatchJSON(s.T(), resp)
	}

}
