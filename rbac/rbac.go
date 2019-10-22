package rbac

import (
	"github.com/casbin/casbin/v2"
)

var rbace *casbin.SyncedEnforcer

// ClearPolicy clears all policy.
func ClearPolicy() {
	rbace.ClearPolicy()
}

// LoadPolicy reloads the policy from file/database.
func LoadPolicy() error {
	return rbace.LoadPolicy()
}

// LoadFilteredPolicy reloads a filtered policy from file/database.
func LoadFilteredPolicy(filter interface{}) error {
	return rbace.LoadFilteredPolicy(filter)
}

// IsFiltered returns true if the loaded policy has been filtered.
func IsFiltered() bool {
	return rbace.IsFiltered()
}

// SavePolicy saves the current policy (usually after changed with Casbin API) back to file/database.
func SavePolicy() error {
	return rbace.SavePolicy()
}

// GetAllSubjects gets the list of subjects that show up in the current policy.
func GetAllSubjects() []string {
	return rbace.GetAllNamedSubjects("p")
}

// GetAllNamedSubjects gets the list of subjects that show up in the current named policy.
func GetAllNamedSubjects(ptype string) []string {
	return rbace.GetAllNamedSubjects(ptype)
}

// GetAllObjects gets the list of objects that show up in the current policy.
func GetAllObjects() []string {
	return rbace.GetAllNamedObjects("p")
}

// GetAllNamedObjects gets the list of objects that show up in the current named policy.
func GetAllNamedObjects(ptype string) []string {
	return rbace.GetAllNamedObjects(ptype)
}

// GetAllActions gets the list of actions that show up in the current policy.
func GetAllActions() []string {
	return rbace.GetAllNamedActions("p")
}

// GetAllNamedActions gets the list of actions that show up in the current named policy.
func GetAllNamedActions(ptype string) []string {
	return rbace.GetAllNamedActions(ptype)
}

// GetAllRoles gets the list of roles that show up in the current policy.
func GetAllRoles() []string {
	return rbace.GetAllNamedRoles("g")
}

// GetAllNamedRoles gets the list of roles that show up in the current named policy.
func GetAllNamedRoles(ptype string) []string {
	return rbace.GetAllNamedRoles(ptype)
}

// GetPolicy gets all the authorization rules in the policy.
func GetPolicy() [][]string {
	return rbace.GetNamedPolicy("p")
}

// GetFilteredPolicy gets all the authorization rules in the policy, field filters can be specified.
func GetFilteredPolicy(fieldIndex int, fieldValues ...string) [][]string {
	return rbace.GetFilteredPolicy(fieldIndex, fieldValues...)
}

// GetNamedPolicy gets all the authorization rules in the named policy.
func GetNamedPolicy(ptype string) [][]string {
	return rbace.GetNamedPolicy(ptype)
}

// GetFilteredNamedPolicy gets all the authorization rules in the named policy, field filters can be specified.
func GetFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) [][]string {
	return rbace.GetFilteredNamedPolicy(ptype, fieldIndex, fieldValues...)
}

// GetGroupingPolicy gets all the role inheritance rules in the policy.
func GetGroupingPolicy() [][]string {
	return rbace.GetNamedGroupingPolicy("g")
}

// GetFilteredGroupingPolicy gets all the role inheritance rules in the policy, field filters can be specified.
func GetFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) [][]string {
	return rbace.GetFilteredNamedGroupingPolicy("g", fieldIndex, fieldValues...)
}

// GetNamedGroupingPolicy gets all the role inheritance rules in the policy.
func GetNamedGroupingPolicy(ptype string) [][]string {
	return rbace.GetNamedGroupingPolicy(ptype)
}

// GetFilteredNamedGroupingPolicy gets all the role inheritance rules in the policy, field filters can be specified.
func GetFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) [][]string {
	return rbace.GetFilteredNamedGroupingPolicy(ptype, fieldIndex, fieldValues...)
}

// HasPolicy determines whether an authorization rule exists.
func HasPolicy(params ...interface{}) bool {
	return rbace.HasNamedPolicy("p", params...)
}

// HasNamedPolicy determines whether a named authorization rule exists.
func HasNamedPolicy(ptype string, params ...interface{}) bool {

	return rbace.HasNamedPolicy(ptype, params...)
}

// AddPolicy adds an authorization rule to the current policy.
// If the rule already exists, the function returns false and the rule will not be added.
// Otherwise the function returns true by adding the new rule.
func AddPolicy(params ...interface{}) (bool, error) {
	return rbace.AddNamedPolicy("p", params...)
}

// AddNamedPolicy adds an authorization rule to the current named policy.
// If the rule already exists, the function returns false and the rule will not be added.
// Otherwise the function returns true by adding the new rule.
func AddNamedPolicy(ptype string, params ...interface{}) (bool, error) {
	return rbace.AddNamedPolicy(ptype, params...)
}

// RemovePolicy removes an authorization rule from the current policy.
func RemovePolicy(params ...interface{}) (bool, error) {
	return rbace.RemoveNamedPolicy("p", params...)
}

// RemoveFilteredPolicy removes an authorization rule from the current policy, field filters can be specified.
func RemoveFilteredPolicy(fieldIndex int, fieldValues ...string) (bool, error) {
	return rbace.RemoveFilteredNamedPolicy("p", fieldIndex, fieldValues...)
}

// RemoveNamedPolicy removes an authorization rule from the current named policy.
func RemoveNamedPolicy(ptype string, params ...interface{}) (bool, error) {
	return rbace.RemoveNamedPolicy(ptype, params...)
}

// RemoveFilteredNamedPolicy removes an authorization rule from the current named policy, field filters can be specified.
func RemoveFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) (bool, error) {
	return rbace.RemoveFilteredNamedPolicy(ptype, fieldIndex, fieldValues...)
}

// HasGroupingPolicy determines whether a role inheritance rule exists.
func HasGroupingPolicy(params ...interface{}) bool {
	return rbace.HasNamedGroupingPolicy("g", params...)
}

// HasNamedGroupingPolicy determines whether a named role inheritance rule exists.
func HasNamedGroupingPolicy(ptype string, params ...interface{}) bool {
	return rbace.HasNamedGroupingPolicy(ptype, params...)
}

// AddGroupingPolicy adds a role inheritance rule to the current policy.
// If the rule already exists, the function returns false and the rule will not be added.
// Otherwise the function returns true by adding the new rule.
func AddGroupingPolicy(params ...interface{}) (bool, error) {
	return rbace.AddNamedGroupingPolicy("g", params...)
}

// AddNamedGroupingPolicy adds a named role inheritance rule to the current policy.
// If the rule already exists, the function returns false and the rule will not be added.
// Otherwise the function returns true by adding the new rule.
func AddNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error) {
	return rbace.AddNamedGroupingPolicy(ptype, params...)
}

// RemoveGroupingPolicy removes a role inheritance rule from the current policy.
func RemoveGroupingPolicy(params ...interface{}) (bool, error) {
	return rbace.RemoveNamedGroupingPolicy("g", params...)
}

// RemoveFilteredGroupingPolicy removes a role inheritance rule from the current policy, field filters can be specified.
func RemoveFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) (bool, error) {
	return rbace.RemoveFilteredNamedGroupingPolicy("g", fieldIndex, fieldValues...)
}

// RemoveNamedGroupingPolicy removes a role inheritance rule from the current named policy.
func RemoveNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error) {
	return rbace.RemoveNamedGroupingPolicy(ptype, params...)
}

// RemoveFilteredNamedGroupingPolicy removes a role inheritance rule from the current named policy, field filters can be specified.
func RemoveFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) (bool, error) {
	return rbace.RemoveFilteredNamedGroupingPolicy(ptype, fieldIndex, fieldValues...)
}

// AddFunction adds a customized function.
func AddFunction(name string, function func(args ...interface{}) (interface{}, error)) {
	rbace.AddFunction(name, function)
}

func GetEnforcer() casbin.SyncedEnforcer {
	return *rbace
}

func registerEnforcer(e *casbin.SyncedEnforcer) error {
	rbace = e
	rbace.EnableAutoSave(true)

	_, err := AddPolicy("admin", ".*", ".*")
	/*
		rbace.AddPolicy(...)
		rbace.RemovePolicy(...)
	*/
	return err
}
