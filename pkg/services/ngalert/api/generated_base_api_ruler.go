/*Package api contains base API implementation of unified alerting
 *
 *Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 *
 *Do not manually edit these files, please find ngalert/api/swagger-codegen/ for commands on how to generate them.
 */

package api

import (
	"net/http"

	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/api/routing"
	"github.com/grafana/grafana/pkg/middleware"
	"github.com/grafana/grafana/pkg/models"
	apimodels "github.com/grafana/grafana/pkg/services/ngalert/api/tooling/definitions"
	"github.com/grafana/grafana/pkg/services/ngalert/metrics"
	"github.com/grafana/grafana/pkg/web"
)

type RulerApiForkingService interface {
	RouteDeleteGrafanaRuleGroupConfig(*models.ReqContext) response.Response
	RouteDeleteNamespaceGrafanaRulesConfig(*models.ReqContext) response.Response
	RouteDeleteNamespaceRulesConfig(*models.ReqContext) response.Response
	RouteDeleteRuleGroupConfig(*models.ReqContext) response.Response
	RouteGetGrafanaRuleGroupConfig(*models.ReqContext) response.Response
	RouteGetGrafanaRulesConfig(*models.ReqContext) response.Response
	RouteGetNamespaceGrafanaRulesConfig(*models.ReqContext) response.Response
	RouteGetNamespaceRulesConfig(*models.ReqContext) response.Response
	RouteGetRulegGroupConfig(*models.ReqContext) response.Response
	RouteGetRulesConfig(*models.ReqContext) response.Response
	RoutePostNameGrafanaRulesConfig(*models.ReqContext) response.Response
	RoutePostNameRulesConfig(*models.ReqContext) response.Response
}

func (f *ForkedRulerApi) RouteDeleteGrafanaRuleGroupConfig(ctx *models.ReqContext) response.Response {
	return f.forkRouteDeleteGrafanaRuleGroupConfig(ctx)
}

func (f *ForkedRulerApi) RouteDeleteNamespaceGrafanaRulesConfig(ctx *models.ReqContext) response.Response {
	return f.forkRouteDeleteNamespaceGrafanaRulesConfig(ctx)
}

func (f *ForkedRulerApi) RouteDeleteNamespaceRulesConfig(ctx *models.ReqContext) response.Response {
	return f.forkRouteDeleteNamespaceRulesConfig(ctx)
}

func (f *ForkedRulerApi) RouteDeleteRuleGroupConfig(ctx *models.ReqContext) response.Response {
	return f.forkRouteDeleteRuleGroupConfig(ctx)
}

func (f *ForkedRulerApi) RouteGetGrafanaRuleGroupConfig(ctx *models.ReqContext) response.Response {
	return f.forkRouteGetGrafanaRuleGroupConfig(ctx)
}

func (f *ForkedRulerApi) RouteGetGrafanaRulesConfig(ctx *models.ReqContext) response.Response {
	return f.forkRouteGetGrafanaRulesConfig(ctx)
}

func (f *ForkedRulerApi) RouteGetNamespaceGrafanaRulesConfig(ctx *models.ReqContext) response.Response {
	return f.forkRouteGetNamespaceGrafanaRulesConfig(ctx)
}

func (f *ForkedRulerApi) RouteGetNamespaceRulesConfig(ctx *models.ReqContext) response.Response {
	return f.forkRouteGetNamespaceRulesConfig(ctx)
}

func (f *ForkedRulerApi) RouteGetRulegGroupConfig(ctx *models.ReqContext) response.Response {
	return f.forkRouteGetRulegGroupConfig(ctx)
}

func (f *ForkedRulerApi) RouteGetRulesConfig(ctx *models.ReqContext) response.Response {
	return f.forkRouteGetRulesConfig(ctx)
}

func (f *ForkedRulerApi) RoutePostNameGrafanaRulesConfig(ctx *models.ReqContext) response.Response {
	conf := apimodels.PostableRuleGroupConfig{}
	if err := web.Bind(ctx.Req, &conf); err != nil {
		return response.Error(http.StatusBadRequest, "bad request data", err)
	}
	return f.forkRoutePostNameGrafanaRulesConfig(ctx, conf)
}

func (f *ForkedRulerApi) RoutePostNameRulesConfig(ctx *models.ReqContext) response.Response {
	conf := apimodels.PostableRuleGroupConfig{}
	if err := web.Bind(ctx.Req, &conf); err != nil {
		return response.Error(http.StatusBadRequest, "bad request data", err)
	}
	return f.forkRoutePostNameRulesConfig(ctx, conf)
}

func (api *API) RegisterRulerApiEndpoints(srv RulerApiForkingService, m *metrics.API) {
	api.RouteRegister.Group("", func(group routing.RouteRegister) {
		group.Delete(
			toMacaronPath("/api/ruler/grafana/api/v1/rules/{Namespace}/{Groupname}"),
			api.authorize(http.MethodDelete, "/api/ruler/grafana/api/v1/rules/{Namespace}/{Groupname}"),
			metrics.Instrument(
				http.MethodDelete,
				"/api/ruler/grafana/api/v1/rules/{Namespace}/{Groupname}",
				srv.RouteDeleteGrafanaRuleGroupConfig,
				m,
			),
		)
		group.Delete(
			toMacaronPath("/api/ruler/grafana/api/v1/rules/{Namespace}"),
			api.authorize(http.MethodDelete, "/api/ruler/grafana/api/v1/rules/{Namespace}"),
			metrics.Instrument(
				http.MethodDelete,
				"/api/ruler/grafana/api/v1/rules/{Namespace}",
				srv.RouteDeleteNamespaceGrafanaRulesConfig,
				m,
			),
		)
		group.Delete(
			toMacaronPath("/api/ruler/{DatasourceUID}/api/v1/rules/{Namespace}"),
			api.authorize(http.MethodDelete, "/api/ruler/{DatasourceUID}/api/v1/rules/{Namespace}"),
			metrics.Instrument(
				http.MethodDelete,
				"/api/ruler/{DatasourceUID}/api/v1/rules/{Namespace}",
				srv.RouteDeleteNamespaceRulesConfig,
				m,
			),
		)
		group.Delete(
			toMacaronPath("/api/ruler/{DatasourceUID}/api/v1/rules/{Namespace}/{Groupname}"),
			api.authorize(http.MethodDelete, "/api/ruler/{DatasourceUID}/api/v1/rules/{Namespace}/{Groupname}"),
			metrics.Instrument(
				http.MethodDelete,
				"/api/ruler/{DatasourceUID}/api/v1/rules/{Namespace}/{Groupname}",
				srv.RouteDeleteRuleGroupConfig,
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/ruler/grafana/api/v1/rules/{Namespace}/{Groupname}"),
			api.authorize(http.MethodGet, "/api/ruler/grafana/api/v1/rules/{Namespace}/{Groupname}"),
			metrics.Instrument(
				http.MethodGet,
				"/api/ruler/grafana/api/v1/rules/{Namespace}/{Groupname}",
				srv.RouteGetGrafanaRuleGroupConfig,
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/ruler/grafana/api/v1/rules"),
			api.authorize(http.MethodGet, "/api/ruler/grafana/api/v1/rules"),
			metrics.Instrument(
				http.MethodGet,
				"/api/ruler/grafana/api/v1/rules",
				srv.RouteGetGrafanaRulesConfig,
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/ruler/grafana/api/v1/rules/{Namespace}"),
			api.authorize(http.MethodGet, "/api/ruler/grafana/api/v1/rules/{Namespace}"),
			metrics.Instrument(
				http.MethodGet,
				"/api/ruler/grafana/api/v1/rules/{Namespace}",
				srv.RouteGetNamespaceGrafanaRulesConfig,
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/ruler/{DatasourceUID}/api/v1/rules/{Namespace}"),
			api.authorize(http.MethodGet, "/api/ruler/{DatasourceUID}/api/v1/rules/{Namespace}"),
			metrics.Instrument(
				http.MethodGet,
				"/api/ruler/{DatasourceUID}/api/v1/rules/{Namespace}",
				srv.RouteGetNamespaceRulesConfig,
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/ruler/{DatasourceUID}/api/v1/rules/{Namespace}/{Groupname}"),
			api.authorize(http.MethodGet, "/api/ruler/{DatasourceUID}/api/v1/rules/{Namespace}/{Groupname}"),
			metrics.Instrument(
				http.MethodGet,
				"/api/ruler/{DatasourceUID}/api/v1/rules/{Namespace}/{Groupname}",
				srv.RouteGetRulegGroupConfig,
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/ruler/{DatasourceUID}/api/v1/rules"),
			api.authorize(http.MethodGet, "/api/ruler/{DatasourceUID}/api/v1/rules"),
			metrics.Instrument(
				http.MethodGet,
				"/api/ruler/{DatasourceUID}/api/v1/rules",
				srv.RouteGetRulesConfig,
				m,
			),
		)
		group.Post(
			toMacaronPath("/api/ruler/grafana/api/v1/rules/{Namespace}"),
			api.authorize(http.MethodPost, "/api/ruler/grafana/api/v1/rules/{Namespace}"),
			metrics.Instrument(
				http.MethodPost,
				"/api/ruler/grafana/api/v1/rules/{Namespace}",
				srv.RoutePostNameGrafanaRulesConfig,
				m,
			),
		)
		group.Post(
			toMacaronPath("/api/ruler/{DatasourceUID}/api/v1/rules/{Namespace}"),
			api.authorize(http.MethodPost, "/api/ruler/{DatasourceUID}/api/v1/rules/{Namespace}"),
			metrics.Instrument(
				http.MethodPost,
				"/api/ruler/{DatasourceUID}/api/v1/rules/{Namespace}",
				srv.RoutePostNameRulesConfig,
				m,
			),
		)
	}, middleware.ReqSignedIn)
}
