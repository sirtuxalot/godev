/* internal/auth/ldap.go */

package auth

import (
  "fmt"
  "github.com/go-ldap/ldap/v3"
  "godev/internal/config"
  "log"
)

// local variables
var app *config.AppConfig
var sizeLimit = 0

// UserLogin struct represents the user credentials.
type UserLogin struct {
  Login_ID  string
  Login_PWD string
}

// connect to ldap server
func Connect() (*ldap.Conn, error) {
  ldapConn, err := ldap.DialURL(app.LDAP_URL)
  if err != nil {
    log.Fatal(fmt.Printf("LDAP connection failed, error details: %v", err))
    return nil, err
  }

  // bind
  if app.LDAP_ID != "" {
    sizeLimit = 1
    // authenticated bind
    err = ldapConn.Bind(app.LDAP_ID, app.LDAP_PWD)
    if err != nil {
      log.Println("LDAP User bind failed:", err)
      return nil, err
    }
    log.Println("***** BAS Domain in use... *****")
  } else {
    // anonymous bind
    err = ldapConn.UnauthenticatedBind("")
    if err != nil {
      log.Println("LDAP anonymous bind failed:", err)
      return nil, err
    }
    log.Println("***** OpenLDAP in use... *****")
  }
  return ldapConn, err
}

// perform ldap authentication
func Auth(ldapConn *ldap.Conn, user UserLogin) (bool, error) {
  // search for the users DN
  userSearch := ldap.NewSearchRequest(app.USER_OU, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, sizeLimit, 0, false, fmt.Sprintf("(%s=%s)", app.ID_LVL, user.Login_ID), []string{"dn"}, nil)
  // search results
  searchResults, err := ldapConn.Search(userSearch)
  if err != nil {
    log.Printf("LDAP search failed for user %s, error details: %v", user.Login_ID, err)
    return false, err
  }
  // ensure there is only one result
  if len(searchResults.Entries) != 1 {
    log.Printf("User: %s not found or multiple entries found", user.Login_ID)
    err = fmt.Errorf("user: %s not found or multiple entries found", user.Login_ID)
    return false, err
  }

  // set user distinguished name
  userDN := searchResults.Entries[0].DN

  // authenticate user
  err = ldapConn.Bind(userDN, user.Login_PWD)
  if err != nil {
    log.Printf("LDAP authentication failed for user %s, error details: %v", user.Login_ID, err)
    err = fmt.Errorf("LDAP authentication failed for user %s", user.Login_ID)
    return false, err
  }
  return true, nil
}
