package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	ory "github.com/ory/kratos-client-go"
)

type AuthParams struct {
	// FlowRedirectURL is the kratos URL to redirect the browser to,
	// when the user wishes to register, and the 'flow' query param is missing
	RegistrationFlowRedirectURL string
	LoginFlowRedirectURL        string
	// VerficationFlowRedirectURL  string
	SettingsFlowRedirectURL string
	RecoveryFlowRedirectURL string
	LogoutFlowRedirectURL   string
}

func (p AuthParams) Registration(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Page: %s\n", "Registration")))
	w.Write([]byte(fmt.Sprintf("Status: %s\n", "To do")))
	w.Write([]byte(fmt.Sprintf("FlowURL: %s", p.RegistrationFlowRedirectURL)))
}

// TODO Better way to instantiate ory kratos client
// TODO Better to get cookie name
// TODO Build UI
func (p AuthParams) Login(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(fmt.Sprintf("Page: %s\n", "Login")))
	// w.Write([]byte(fmt.Sprintf("Status: %s\n", "To do")))
	// w.Write([]byte(fmt.Sprintf("FlowURL: %s", p.LoginFlowRedirectURL)))

	log.Println("Login page")

	// Start the login flow with Kratos if required
	flowID := r.URL.Query().Get("flow")
	if flowID == "" {
		log.Printf("- No flow ID found in URL, initializing login flow, redirect to %s", p.LoginFlowRedirectURL)
		http.Redirect(w, r, p.LoginFlowRedirectURL, http.StatusMovedPermanently)
		return
	} else {
		log.Printf("- Have flow ID: %s\n", flowID)
	}

	configuration := ory.NewConfiguration()
	configuration.Servers = []ory.ServerConfiguration{
		{
			URL: "http://127.0.0.1:4433", // Kratos Admin API
		},
	}

	client := ory.NewAPIClient(configuration)

	cn := "csrf_token_82b119fa58a0a1cb6faa9738c1d0dbbf04fcc89a657b7beb31fcde400ced48ab"
	cookie, err := r.Cookie(cn)
	if err != nil {
		log.Printf("- No cookie found for name %s\n", cn)
	} else {
		log.Printf("- Cookie %s\n", cookie.String())
		log.Printf("- Cookie name %s\n", cookie.Name)
		log.Printf("- Cookie value %s\n", cookie.Value)
	}

	// loginFlow, _, err := api_client.PublicClient().FrontendApi.GetLoginFlow(r.Context()).Id(flowID).Cookie(cookie.String()).Execute()
	flow, _, err := client.FrontendApi.GetLoginFlow(r.Context()).Id(flowID).Cookie(cookie.String()).Execute()
	// loginFlow, res, err := api_client.PublicClient().FrontendApi.CreateBrowserLoginFlow(context.Background()).Execute()c
	// log.Printf("- Create login flow response: %v\n", res.Status)
	if err != nil {
		log.Printf("- Create login flow error: %v\n", err)
	} else {
		log.Println("- Created login flow")
		printJSONPretty(flow)
	}
	// TODO check if the flow has any error messages
	// TODO add logout link
	// TODO add registration query
	// TODO add registration link
	// TODO add recovery link
	// if (flow.ui.messages && flow.ui.messages.length > 0) {
	// 	// the login requires that the user verifies their email address before logging in
	// 	if (flow.ui.messages.some(({ id }) => id === 4000010)) {
	// 	  // we will create a new verification flow and redirect the user to the verification page
	// 	  return redirectToVerificationFlow(flow)
	// 	}
	// }

	// config := api_client.PublicClient().GetConfig().Host
	// fmt.Printf("client config: %v\n", config)

	// api_client.AdminClient().FrontendApi.ListMySessions(context.Background()).Cookie()

	// loginFlowReq := api_client.PublicClient().FrontendApi.GetLoginFlow(r.Context()).Id(flowID)
	// fmt.Printf("- loginFlowReq: %+v\n", loginFlowReq)

	// // loginFlow, _, err := api_client.PublicClient().FrontendApi.GetLoginFlow(r.Context()).Id(flowID).Execute()
	// loginFlow, _, err := loginFlowReq.Execute()
	// // loginFlow, res, err := api_client.PublicClient().FrontendApi.CreateBrowserLoginFlow(context.Background()).Execute()
	// // log.Printf("- Create login flow response: %v\n", res.Status)
	// if err != nil {
	// 	log.Printf("- Create login flow error: %v\n", err)
	// } else {
	// 	log.Printf("- Created login flow: %v\n", loginFlow)
	// }

	// fmt.Printf("login res: %v\n", res)

	// res := api_client.PublicClient().FrontendApi.GetLoginFlow(context.Background())
	// fmt.Printf("login res: %v\n", res)
	// res.get

	// Call Kratos to retrieve the login form
	// params := kratos.NewGetSelfServiceLoginFlowParams

	// Call Kratos to retrieve the login form
	// params := public.NewGetSelfServiceLoginFlowParams()
	// params.SetID(flow)
	// log.Print("Calling Kratos API to get self service login")
	// res, err := api_client.PublicClient().Public.GetSelfServiceLoginFlow(params)
	// if err != nil {
	// 	log.Printf("Error getting self service login flow: %v, redirecting to /", err)
	// 	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	// 	return
	// }
	// dataMap := map[string]interface{}{
	// 	"flow":        flow,
	// 	"config":      res.GetPayload().Methods["password"].Config,
	// 	"fs":          lp.FS,
	// 	"pageHeading": "Login",
	// }
	// if err = GetTemplate(loginPage).Render("layout", w, r, dataMap); err != nil {
	// 	ErrorHandler(w, r, err)
	// }
}

func (p AuthParams) Settings(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Page: %s\n", "Settings")))
	w.Write([]byte(fmt.Sprintf("Status: %s\n", "To do")))
	w.Write([]byte(fmt.Sprintf("FlowURL: %s", p.SettingsFlowRedirectURL)))
}

func (p AuthParams) Recovery(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Page: %s\n", "Recovery")))
	w.Write([]byte(fmt.Sprintf("Status: %s\n", "To do")))
	w.Write([]byte(fmt.Sprintf("FlowURL: %s", p.RecoveryFlowRedirectURL)))
}

func (p AuthParams) Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Page: %s\n", "Logout")))
	w.Write([]byte(fmt.Sprintf("Status: %s\n", "To do")))
	w.Write([]byte(fmt.Sprintf("FlowURL: %s", p.LogoutFlowRedirectURL)))
}

func printJSONPretty(v interface{}) {
	out, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(out))
}
