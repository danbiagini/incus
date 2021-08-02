//go:build linux && cgo && !agent
// +build linux,cgo,!agent

package db

// The code below was generated by lxd-generate - DO NOT EDIT!

import (
	"database/sql"
	"fmt"
	"github.com/lxc/lxd/lxd/db/cluster"
	"github.com/lxc/lxd/lxd/db/query"
	"github.com/lxc/lxd/shared/api"
	"github.com/pkg/errors"
)

var _ = api.ServerEnvironment{}

var instanceObjects = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  ORDER BY projects.id, instances.name
`)

var instanceObjectsByProject = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndType = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.type = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndTypeAndNode = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.type = ? AND node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndTypeAndNodeAndName = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.type = ? AND node = ? AND instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndTypeAndName = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.type = ? AND instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndName = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndNameAndNode = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND instances.name = ? AND node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByProjectAndNode = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE project = ? AND node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByType = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE instances.type = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByTypeAndName = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE instances.type = ? AND instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByTypeAndNameAndNode = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE instances.type = ? AND instances.name = ? AND node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByTypeAndNode = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE instances.type = ? AND node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByNode = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE node = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByNodeAndName = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE node = ? AND instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceObjectsByName = cluster.RegisterStmt(`
SELECT instances.id, projects.name AS project, instances.name, nodes.name AS node, instances.type, instances.architecture, instances.ephemeral, instances.creation_date, instances.stateful, instances.last_use_date, coalesce(instances.description, ''), instances.expiry_date
  FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE instances.name = ? ORDER BY projects.id, instances.name
`)

var instanceProfilesRef = cluster.RegisterStmt(`
SELECT project, name, value FROM instances_profiles_ref ORDER BY project, name
`)

var instanceProfilesRefByProject = cluster.RegisterStmt(`
SELECT project, name, value FROM instances_profiles_ref WHERE project = ? ORDER BY project, name
`)

var instanceProfilesRefByNode = cluster.RegisterStmt(`
SELECT project, name, value FROM instances_profiles_ref WHERE node = ? ORDER BY project, name
`)

var instanceProfilesRefByProjectAndNode = cluster.RegisterStmt(`
SELECT project, name, value FROM instances_profiles_ref WHERE project = ? AND node = ? ORDER BY project, name
`)

var instanceProfilesRefByProjectAndName = cluster.RegisterStmt(`
SELECT project, name, value FROM instances_profiles_ref WHERE project = ? AND name = ? ORDER BY project, name
`)

var instanceConfigRef = cluster.RegisterStmt(`
SELECT project, name, key, value FROM instances_config_ref ORDER BY project, name
`)

var instanceConfigRefByProject = cluster.RegisterStmt(`
SELECT project, name, key, value FROM instances_config_ref WHERE project = ? ORDER BY project, name
`)

var instanceConfigRefByNode = cluster.RegisterStmt(`
SELECT project, name, key, value FROM instances_config_ref WHERE node = ? ORDER BY project, name
`)

var instanceConfigRefByProjectAndNode = cluster.RegisterStmt(`
SELECT project, name, key, value FROM instances_config_ref WHERE project = ? AND node = ? ORDER BY project, name
`)

var instanceConfigRefByProjectAndName = cluster.RegisterStmt(`
SELECT project, name, key, value FROM instances_config_ref WHERE project = ? AND name = ? ORDER BY project, name
`)

var instanceDevicesRef = cluster.RegisterStmt(`
SELECT project, name, device, type, key, value FROM instances_devices_ref ORDER BY project, name
`)

var instanceDevicesRefByProject = cluster.RegisterStmt(`
SELECT project, name, device, type, key, value FROM instances_devices_ref WHERE project = ? ORDER BY project, name
`)

var instanceDevicesRefByNode = cluster.RegisterStmt(`
SELECT project, name, device, type, key, value FROM instances_devices_ref WHERE node = ? ORDER BY project, name
`)

var instanceDevicesRefByProjectAndNode = cluster.RegisterStmt(`
SELECT project, name, device, type, key, value FROM instances_devices_ref WHERE project = ? AND node = ? ORDER BY project, name
`)

var instanceDevicesRefByProjectAndName = cluster.RegisterStmt(`
SELECT project, name, device, type, key, value FROM instances_devices_ref WHERE project = ? AND name = ? ORDER BY project, name
`)

var instanceID = cluster.RegisterStmt(`
SELECT instances.id FROM instances JOIN projects ON instances.project_id = projects.id JOIN nodes ON instances.node_id = nodes.id
  WHERE projects.name = ? AND instances.name = ?
`)

var instanceCreate = cluster.RegisterStmt(`
INSERT INTO instances (project_id, name, node_id, type, architecture, ephemeral, creation_date, stateful, last_use_date, description, expiry_date)
  VALUES ((SELECT projects.id FROM projects WHERE projects.name = ?), ?, (SELECT nodes.id FROM nodes WHERE nodes.name = ?), ?, ?, ?, ?, ?, ?, ?, ?)
`)

var instanceCreateConfigRef = cluster.RegisterStmt(`
INSERT INTO instances_config (instance_id, key, value)
  VALUES (?, ?, ?)
`)

var instanceCreateDevicesRef = cluster.RegisterStmt(`
INSERT INTO instances_devices (instance_id, name, type)
  VALUES (?, ?, ?)
`)
var instanceCreateDevicesConfigRef = cluster.RegisterStmt(`
INSERT INTO instances_devices_config (instance_device_id, key, value)
  VALUES (?, ?, ?)
`)

var instanceRename = cluster.RegisterStmt(`
UPDATE instances SET name = ? WHERE project_id = (SELECT projects.id FROM projects WHERE projects.name = ?) AND name = ?
`)

var instanceDeleteByProjectAndName = cluster.RegisterStmt(`
DELETE FROM instances WHERE project_id = (SELECT projects.id FROM projects WHERE projects.name = ?) AND name = ?
`)

var instanceDeleteConfigRef = cluster.RegisterStmt(`
DELETE FROM instances_config WHERE instance_id = ?
`)

var instanceDeleteDevicesRef = cluster.RegisterStmt(`
DELETE FROM instances_devices WHERE instance_id = ?
`)

var instanceDeleteProfilesRef = cluster.RegisterStmt(`
DELETE FROM instances_profiles WHERE instance_id = ?
`)

var instanceUpdate = cluster.RegisterStmt(`
UPDATE instances
  SET project_id = (SELECT id FROM projects WHERE name = ?), name = ?, node_id = (SELECT id FROM nodes WHERE name = ?), type = ?, architecture = ?, ephemeral = ?, creation_date = ?, stateful = ?, last_use_date = ?, description = ?, expiry_date = ?
 WHERE id = ?
`)

// GetInstances returns all available instances.
// generator: instance GetMany
func (c *ClusterTx) GetInstances(filter InstanceFilter) ([]Instance, error) {
	// Result slice.
	objects := make([]Instance, 0)

	// Check which filter criteria are active.
	criteria := map[string]interface{}{}
	if filter.Project != "" {
		criteria["Project"] = filter.Project
	}
	if filter.Name != "" {
		criteria["Name"] = filter.Name
	}
	if filter.Node != "" {
		criteria["Node"] = filter.Node
	}
	if filter.Type != -1 {
		criteria["Type"] = filter.Type
	}

	// Pick the prepared statement and arguments to use based on active criteria.
	var stmt *sql.Stmt
	var args []interface{}

	if criteria["Project"] != nil && criteria["Type"] != nil && criteria["Node"] != nil && criteria["Name"] != nil {
		stmt = c.stmt(instanceObjectsByProjectAndTypeAndNodeAndName)
		args = []interface{}{
			filter.Project,
			filter.Type,
			filter.Node,
			filter.Name,
		}
	} else if criteria["Project"] != nil && criteria["Type"] != nil && criteria["Node"] != nil {
		stmt = c.stmt(instanceObjectsByProjectAndTypeAndNode)
		args = []interface{}{
			filter.Project,
			filter.Type,
			filter.Node,
		}
	} else if criteria["Project"] != nil && criteria["Type"] != nil && criteria["Name"] != nil {
		stmt = c.stmt(instanceObjectsByProjectAndTypeAndName)
		args = []interface{}{
			filter.Project,
			filter.Type,
			filter.Name,
		}
	} else if criteria["Type"] != nil && criteria["Name"] != nil && criteria["Node"] != nil {
		stmt = c.stmt(instanceObjectsByTypeAndNameAndNode)
		args = []interface{}{
			filter.Type,
			filter.Name,
			filter.Node,
		}
	} else if criteria["Project"] != nil && criteria["Name"] != nil && criteria["Node"] != nil {
		stmt = c.stmt(instanceObjectsByProjectAndNameAndNode)
		args = []interface{}{
			filter.Project,
			filter.Name,
			filter.Node,
		}
	} else if criteria["Project"] != nil && criteria["Type"] != nil {
		stmt = c.stmt(instanceObjectsByProjectAndType)
		args = []interface{}{
			filter.Project,
			filter.Type,
		}
	} else if criteria["Type"] != nil && criteria["Node"] != nil {
		stmt = c.stmt(instanceObjectsByTypeAndNode)
		args = []interface{}{
			filter.Type,
			filter.Node,
		}
	} else if criteria["Type"] != nil && criteria["Name"] != nil {
		stmt = c.stmt(instanceObjectsByTypeAndName)
		args = []interface{}{
			filter.Type,
			filter.Name,
		}
	} else if criteria["Project"] != nil && criteria["Node"] != nil {
		stmt = c.stmt(instanceObjectsByProjectAndNode)
		args = []interface{}{
			filter.Project,
			filter.Node,
		}
	} else if criteria["Project"] != nil && criteria["Name"] != nil {
		stmt = c.stmt(instanceObjectsByProjectAndName)
		args = []interface{}{
			filter.Project,
			filter.Name,
		}
	} else if criteria["Node"] != nil && criteria["Name"] != nil {
		stmt = c.stmt(instanceObjectsByNodeAndName)
		args = []interface{}{
			filter.Node,
			filter.Name,
		}
	} else if criteria["Type"] != nil {
		stmt = c.stmt(instanceObjectsByType)
		args = []interface{}{
			filter.Type,
		}
	} else if criteria["Project"] != nil {
		stmt = c.stmt(instanceObjectsByProject)
		args = []interface{}{
			filter.Project,
		}
	} else if criteria["Node"] != nil {
		stmt = c.stmt(instanceObjectsByNode)
		args = []interface{}{
			filter.Node,
		}
	} else if criteria["Name"] != nil {
		stmt = c.stmt(instanceObjectsByName)
		args = []interface{}{
			filter.Name,
		}
	} else {
		stmt = c.stmt(instanceObjects)
		args = []interface{}{}
	}

	// Dest function for scanning a row.
	dest := func(i int) []interface{} {
		objects = append(objects, Instance{})
		return []interface{}{
			&objects[i].ID,
			&objects[i].Project,
			&objects[i].Name,
			&objects[i].Node,
			&objects[i].Type,
			&objects[i].Architecture,
			&objects[i].Ephemeral,
			&objects[i].CreationDate,
			&objects[i].Stateful,
			&objects[i].LastUseDate,
			&objects[i].Description,
			&objects[i].ExpiryDate,
		}
	}

	// Select.
	err := query.SelectObjects(stmt, dest, args...)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to fetch instances")
	}

	// Fill field Config.
	configObjects, err := c.InstanceConfigRef(filter)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to fetch field Config")
	}

	for i := range objects {
		_, ok0 := configObjects[objects[i].Project]
		if !ok0 {
			subIndex := map[string]map[string]string{}
			configObjects[objects[i].Project] = subIndex
		}

		value := configObjects[objects[i].Project][objects[i].Name]
		if value == nil {
			value = map[string]string{}
		}
		objects[i].Config = value
	}

	// Fill field Devices.
	devicesObjects, err := c.InstanceDevicesRef(filter)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to fetch field Devices")
	}

	for i := range objects {
		_, ok0 := devicesObjects[objects[i].Project]
		if !ok0 {
			subIndex := map[string]map[string]map[string]string{}
			devicesObjects[objects[i].Project] = subIndex
		}

		value := devicesObjects[objects[i].Project][objects[i].Name]
		if value == nil {
			value = map[string]map[string]string{}
		}
		objects[i].Devices = value
	}

	// Fill field Profiles.
	profilesObjects, err := c.InstanceProfilesRef(filter)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to fetch field Profiles")
	}

	for i := range objects {
		_, ok0 := profilesObjects[objects[i].Project]
		if !ok0 {
			subIndex := map[string][]string{}
			profilesObjects[objects[i].Project] = subIndex
		}

		value := profilesObjects[objects[i].Project][objects[i].Name]
		if value == nil {
			value = []string{}
		}
		objects[i].Profiles = value
	}

	return objects, nil
}

// GetInstance returns the instance with the given key.
// generator: instance GetOne
func (c *ClusterTx) GetInstance(project string, name string) (*Instance, error) {
	filter := InstanceFilter{}
	filter.Project = project
	filter.Name = name
	filter.Type = -1

	objects, err := c.GetInstances(filter)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to fetch Instance")
	}

	switch len(objects) {
	case 0:
		return nil, ErrNoSuchObject
	case 1:
		return &objects[0], nil
	default:
		return nil, fmt.Errorf("More than one instance matches")
	}
}

// GetInstanceID return the ID of the instance with the given key.
// generator: instance ID
func (c *ClusterTx) GetInstanceID(project string, name string) (int64, error) {
	stmt := c.stmt(instanceID)
	rows, err := stmt.Query(project, name)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to get instance ID")
	}
	defer rows.Close()

	// Ensure we read one and only one row.
	if !rows.Next() {
		return -1, ErrNoSuchObject
	}
	var id int64
	err = rows.Scan(&id)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to scan ID")
	}
	if rows.Next() {
		return -1, fmt.Errorf("More than one row returned")
	}
	err = rows.Err()
	if err != nil {
		return -1, errors.Wrap(err, "Result set failure")
	}

	return id, nil
}

// InstanceExists checks if a instance with the given key exists.
// generator: instance Exists
func (c *ClusterTx) InstanceExists(project string, name string) (bool, error) {
	_, err := c.GetInstanceID(project, name)
	if err != nil {
		if err == ErrNoSuchObject {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// CreateInstance adds a new instance to the database.
// generator: instance Create
func (c *ClusterTx) CreateInstance(object Instance) (int64, error) {
	// Check if a instance with the same key exists.
	exists, err := c.InstanceExists(object.Project, object.Name)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to check for duplicates")
	}
	if exists {
		return -1, fmt.Errorf("This instance already exists")
	}

	args := make([]interface{}, 11)

	// Populate the statement arguments.
	args[0] = object.Project
	args[1] = object.Name
	args[2] = object.Node
	args[3] = object.Type
	args[4] = object.Architecture
	args[5] = object.Ephemeral
	args[6] = object.CreationDate
	args[7] = object.Stateful
	args[8] = object.LastUseDate
	args[9] = object.Description
	args[10] = object.ExpiryDate

	// Prepared statement to use.
	stmt := c.stmt(instanceCreate)

	// Execute the statement.
	result, err := stmt.Exec(args...)
	if err != nil {
		return -1, errors.Wrap(err, "Failed to create instance")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, errors.Wrap(err, "Failed to fetch instance ID")
	}

	// Insert config reference.
	stmt = c.stmt(instanceCreateConfigRef)
	for key, value := range object.Config {
		_, err := stmt.Exec(id, key, value)
		if err != nil {
			return -1, errors.Wrap(err, "Insert config for instance")
		}
	}

	// Insert devices reference.
	for name, config := range object.Devices {
		typ, ok := config["type"]
		if !ok {
			return -1, fmt.Errorf("No type for device %s", name)
		}
		typCode, err := deviceTypeToInt(typ)
		if err != nil {
			return -1, errors.Wrapf(err, "Device type code for %s", typ)
		}
		stmt = c.stmt(instanceCreateDevicesRef)
		result, err := stmt.Exec(id, name, typCode)
		if err != nil {
			return -1, errors.Wrapf(err, "Insert device %s", name)
		}
		deviceID, err := result.LastInsertId()
		if err != nil {
			return -1, errors.Wrap(err, "Failed to fetch device ID")
		}
		stmt = c.stmt(instanceCreateDevicesConfigRef)
		for key, value := range config {
			_, err := stmt.Exec(deviceID, key, value)
			if err != nil {
				return -1, errors.Wrap(err, "Insert config for instance")
			}
		}
	}

	// Insert profiles reference.
	err = addProfilesToInstance(c.tx, int(id), object.Project, object.Profiles)
	if err != nil {
		return -1, errors.Wrap(err, "Insert profiles for instance")
	}
	return id, nil
}

// InstanceProfilesRef returns entities used by instances.
// generator: instance ProfilesRef
func (c *ClusterTx) InstanceProfilesRef(filter InstanceFilter) (map[string]map[string][]string, error) {
	// Result slice.
	objects := make([]struct {
		Project string
		Name    string
		Value   string
	}, 0)

	// Check which filter criteria are active.
	criteria := map[string]interface{}{}
	if filter.Project != "" {
		criteria["Project"] = filter.Project
	}
	if filter.Name != "" {
		criteria["Name"] = filter.Name
	}

	// Pick the prepared statement and arguments to use based on active criteria.
	var stmt *sql.Stmt
	var args []interface{}

	if criteria["Project"] != nil && criteria["Node"] != nil {
		stmt = c.stmt(instanceProfilesRefByProjectAndNode)
		args = []interface{}{
			filter.Project,
			filter.Node,
		}
	} else if criteria["Project"] != nil && criteria["Name"] != nil {
		stmt = c.stmt(instanceProfilesRefByProjectAndName)
		args = []interface{}{
			filter.Project,
			filter.Name,
		}
	} else if criteria["Project"] != nil {
		stmt = c.stmt(instanceProfilesRefByProject)
		args = []interface{}{
			filter.Project,
		}
	} else if criteria["Node"] != nil {
		stmt = c.stmt(instanceProfilesRefByNode)
		args = []interface{}{
			filter.Node,
		}
	} else {
		stmt = c.stmt(instanceProfilesRef)
		args = []interface{}{}
	}

	// Dest function for scanning a row.
	dest := func(i int) []interface{} {
		objects = append(objects, struct {
			Project string
			Name    string
			Value   string
		}{})
		return []interface{}{
			&objects[i].Project,
			&objects[i].Name,
			&objects[i].Value,
		}
	}

	// Select.
	err := query.SelectObjects(stmt, dest, args...)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to fetch string ref for instances")
	}

	// Build index by primary name.
	index := map[string]map[string][]string{}

	for _, object := range objects {
		_, ok0 := index[object.Project]
		if !ok0 {
			subIndex := map[string][]string{}
			index[object.Project] = subIndex
		}

		item, ok := index[object.Project][object.Name]
		if !ok {
			item = []string{}
		}

		index[object.Project][object.Name] = append(item, object.Value)
	}

	return index, nil
}

// InstanceConfigRef returns entities used by instances.
// generator: instance ConfigRef
func (c *ClusterTx) InstanceConfigRef(filter InstanceFilter) (map[string]map[string]map[string]string, error) {
	// Result slice.
	objects := make([]struct {
		Project string
		Name    string
		Key     string
		Value   string
	}, 0)

	// Check which filter criteria are active.
	criteria := map[string]interface{}{}
	if filter.Project != "" {
		criteria["Project"] = filter.Project
	}
	if filter.Name != "" {
		criteria["Name"] = filter.Name
	}

	// Pick the prepared statement and arguments to use based on active criteria.
	var stmt *sql.Stmt
	var args []interface{}

	if criteria["Project"] != nil && criteria["Node"] != nil {
		stmt = c.stmt(instanceConfigRefByProjectAndNode)
		args = []interface{}{
			filter.Project,
			filter.Node,
		}
	} else if criteria["Project"] != nil && criteria["Name"] != nil {
		stmt = c.stmt(instanceConfigRefByProjectAndName)
		args = []interface{}{
			filter.Project,
			filter.Name,
		}
	} else if criteria["Project"] != nil {
		stmt = c.stmt(instanceConfigRefByProject)
		args = []interface{}{
			filter.Project,
		}
	} else if criteria["Node"] != nil {
		stmt = c.stmt(instanceConfigRefByNode)
		args = []interface{}{
			filter.Node,
		}
	} else {
		stmt = c.stmt(instanceConfigRef)
		args = []interface{}{}
	}

	// Dest function for scanning a row.
	dest := func(i int) []interface{} {
		objects = append(objects, struct {
			Project string
			Name    string
			Key     string
			Value   string
		}{})
		return []interface{}{
			&objects[i].Project,
			&objects[i].Name,
			&objects[i].Key,
			&objects[i].Value,
		}
	}

	// Select.
	err := query.SelectObjects(stmt, dest, args...)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to fetch  ref for instances")
	}

	// Build index by primary name.
	index := map[string]map[string]map[string]string{}

	for _, object := range objects {
		_, ok0 := index[object.Project]
		if !ok0 {
			subIndex := map[string]map[string]string{}
			index[object.Project] = subIndex
		}

		item, ok := index[object.Project][object.Name]
		if !ok {
			item = map[string]string{}
		}

		index[object.Project][object.Name] = item
		item[object.Key] = object.Value
	}

	return index, nil
}

// InstanceDevicesRef returns entities used by instances.
// generator: instance DevicesRef
func (c *ClusterTx) InstanceDevicesRef(filter InstanceFilter) (map[string]map[string]map[string]map[string]string, error) {
	// Result slice.
	objects := make([]struct {
		Project string
		Name    string
		Device  string
		Type    int
		Key     string
		Value   string
	}, 0)

	// Check which filter criteria are active.
	criteria := map[string]interface{}{}
	if filter.Project != "" {
		criteria["Project"] = filter.Project
	}
	if filter.Name != "" {
		criteria["Name"] = filter.Name
	}

	// Pick the prepared statement and arguments to use based on active criteria.
	var stmt *sql.Stmt
	var args []interface{}

	if criteria["Project"] != nil && criteria["Node"] != nil {
		stmt = c.stmt(instanceDevicesRefByProjectAndNode)
		args = []interface{}{
			filter.Project,
			filter.Node,
		}
	} else if criteria["Project"] != nil && criteria["Name"] != nil {
		stmt = c.stmt(instanceDevicesRefByProjectAndName)
		args = []interface{}{
			filter.Project,
			filter.Name,
		}
	} else if criteria["Project"] != nil {
		stmt = c.stmt(instanceDevicesRefByProject)
		args = []interface{}{
			filter.Project,
		}
	} else if criteria["Node"] != nil {
		stmt = c.stmt(instanceDevicesRefByNode)
		args = []interface{}{
			filter.Node,
		}
	} else {
		stmt = c.stmt(instanceDevicesRef)
		args = []interface{}{}
	}

	// Dest function for scanning a row.
	dest := func(i int) []interface{} {
		objects = append(objects, struct {
			Project string
			Name    string
			Device  string
			Type    int
			Key     string
			Value   string
		}{})
		return []interface{}{
			&objects[i].Project,
			&objects[i].Name,
			&objects[i].Device,
			&objects[i].Type,
			&objects[i].Key,
			&objects[i].Value,
		}
	}

	// Select.
	err := query.SelectObjects(stmt, dest, args...)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to fetch  ref for instances")
	}

	// Build index by primary name.
	index := map[string]map[string]map[string]map[string]string{}

	for _, object := range objects {
		_, ok0 := index[object.Project]
		if !ok0 {
			subIndex := map[string]map[string]map[string]string{}
			index[object.Project] = subIndex
		}

		item, ok := index[object.Project][object.Name]
		if !ok {
			item = map[string]map[string]string{}
		}

		index[object.Project][object.Name] = item
		config, ok := item[object.Device]
		if !ok {
			// First time we see this device, let's int the config
			// and add the type.
			deviceType, err := deviceTypeToString(object.Type)
			if err != nil {
				return nil, errors.Wrapf(
					err, "unexpected device type code '%d'", object.Type)
			}
			config = map[string]string{}
			config["type"] = deviceType
			item[object.Device] = config
		}
		if object.Key != "" {
			config[object.Key] = object.Value
		}
	}

	return index, nil
}

// RenameInstance renames the instance matching the given key parameters.
// generator: instance Rename
func (c *ClusterTx) RenameInstance(project string, name string, to string) error {
	stmt := c.stmt(instanceRename)
	result, err := stmt.Exec(to, project, name)
	if err != nil {
		return errors.Wrap(err, "Rename instance")
	}

	n, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Fetch affected rows")
	}
	if n != 1 {
		return fmt.Errorf("Query affected %d rows instead of 1", n)
	}
	return nil
}

// DeleteInstance deletes the instance matching the given key parameters.
// generator: instance DeleteOne-by-Project-and-Name
func (c *ClusterTx) DeleteInstance(project string, name string) error {
	stmt := c.stmt(instanceDeleteByProjectAndName)
	result, err := stmt.Exec(project, name)
	if err != nil {
		return errors.Wrap(err, "Delete instance")
	}

	n, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Fetch affected rows")
	}
	if n != 1 {
		return fmt.Errorf("Query deleted %d rows instead of 1", n)
	}

	return nil
}

// UpdateInstance updates the instance matching the given key parameters.
// generator: instance Update
func (c *ClusterTx) UpdateInstance(project string, name string, object Instance) error {
	id, err := c.GetInstanceID(project, name)
	if err != nil {
		return errors.Wrap(err, "Get instance")
	}

	stmt := c.stmt(instanceUpdate)
	result, err := stmt.Exec(object.Project, object.Name, object.Node, object.Type, object.Architecture, object.Ephemeral, object.CreationDate, object.Stateful, object.LastUseDate, object.Description, object.ExpiryDate, id)
	if err != nil {
		return errors.Wrap(err, "Update instance")
	}

	n, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Fetch affected rows")
	}
	if n != 1 {
		return fmt.Errorf("Query updated %d rows instead of 1", n)
	}

	// Delete current config.
	stmt = c.stmt(instanceDeleteConfigRef)
	_, err = stmt.Exec(id)
	if err != nil {
		return errors.Wrap(err, "Delete current config")
	}

	// Insert config reference.
	stmt = c.stmt(instanceCreateConfigRef)
	for key, value := range object.Config {
		if value == "" {
			continue
		}
		_, err := stmt.Exec(id, key, value)
		if err != nil {
			return errors.Wrap(err, "Insert config for instance")
		}
	}

	// Delete current devices.
	stmt = c.stmt(instanceDeleteDevicesRef)
	_, err = stmt.Exec(id)
	if err != nil {
		return errors.Wrap(err, "Delete current devices")
	}

	// Insert devices reference.
	for name, config := range object.Devices {
		typ, ok := config["type"]
		if !ok {
			return fmt.Errorf("No type for device %s", name)
		}
		typCode, err := deviceTypeToInt(typ)
		if err != nil {
			return errors.Wrapf(err, "Device type code for %s", typ)
		}
		stmt = c.stmt(instanceCreateDevicesRef)
		result, err := stmt.Exec(id, name, typCode)
		if err != nil {
			return errors.Wrapf(err, "Insert device %s", name)
		}
		deviceID, err := result.LastInsertId()
		if err != nil {
			return errors.Wrap(err, "Failed to fetch device ID")
		}
		stmt = c.stmt(instanceCreateDevicesConfigRef)
		for key, value := range config {
			if value == "" {
				continue
			}
			_, err := stmt.Exec(deviceID, key, value)
			if err != nil {
				return errors.Wrap(err, "Insert config for instance")
			}
		}
	}

	// Delete current profiles.
	stmt = c.stmt(instanceDeleteProfilesRef)
	_, err = stmt.Exec(id)
	if err != nil {
		return errors.Wrap(err, "Delete current profiles")
	}

	// Insert profiles reference.
	err = addProfilesToInstance(c.tx, int(id), object.Project, object.Profiles)
	if err != nil {
		return errors.Wrap(err, "Insert profiles for instance")
	}
	return nil
}
