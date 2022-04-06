package controllers

// Login, cria o logon para a api
/*func (c *Controller) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.RespondError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var account accounts.Account
	if err := json.Unmarshal(bodyRequest, &account); err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	defer r.Body.Close()

	modelFindDocument := db.Database{}
	verifyDocument, err := modelFindDocument.FindDocument()
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	err = security.VerifyPasswd(verifyDocument.Secret, account.Secret)
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	token, err := security.CreateToken(verifyDocument.ID)
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
	}

	responses.RespondJSON(w, http.StatusOK, models.Authentication{ID: verifyDocument.ID, Token: token})
}*/
