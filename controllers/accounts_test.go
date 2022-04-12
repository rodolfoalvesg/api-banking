package controllers

// TestCreateAccount, teste do handler de criação de conta

/*func TestCreateAccount(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name string
		accountMock
	}

}*/

/*

//TestHandlerShowBalance, teste do handler para exibição de saldo
func TestHandlerShowBalance(t *testing.T) {
	t.Parallel()

	accFake := models.Account{
		ID:     "dfsh15hjfg4hgfsdhgdsf",
		Secret: "123456789",
	}

	accListA, _ := accounts.NewCreateAccount(accFake)
	accListB := models.Account{}

	controller := NewController(nil)

	testShowBalance := map[string]struct {
		accBalanceID models.Account
		want         int
	}{
		"Status OK":  {accListA, http.StatusOK},
		"Status Bad": {accListB, http.StatusBadRequest},
	}

	for name, tt := range testShowBalance {

		path := fmt.Sprintf("/accounts/%s/balance", tt.accBalanceID.ID)

		request := httptest.NewRequest(http.MethodGet, path, nil)
		response := httptest.NewRecorder()

		vars := map[string]string{
			"account_id": tt.accBalanceID.ID,
		}

		request = mux.SetURLVars(request, vars)

		controller.HandlerShowBalance(response, request)

		respondeCode := response.Result().StatusCode

		if respondeCode != tt.want {
			t.Errorf("%s: got %v, want %v", name, respondeCode, tt.want)
		}

	}
}

// TestHandlerShowAccounts, teste do handler para listagem de conta
func TestHandlerShowAccounts(t *testing.T) {
	t.Parallel()

	controller := NewController(nil)

	request := httptest.NewRequest(http.MethodGet, "/accounts", nil)
	response := httptest.NewRecorder()

	controller.HandlerShowAccounts(response, request)

	fmt.Println(response)

	if response.Result().StatusCode != http.StatusOK {
		t.Fail()
	}
}
*/
