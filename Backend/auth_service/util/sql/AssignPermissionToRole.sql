-- Get role and permission IDs by name
DO $$
DECLARE
    roleName VARCHAR(255) := 'ADMIN';
    permissionName VARCHAR(255) := 'LoggedInSO_ru';
    roleId UUID;
    permissionId UUID;
BEGIN
    SELECT id INTO roleId
    FROM roles
    WHERE name = roleName;

    SELECT id INTO permissionId
    FROM permissions
    WHERE name = permissionName;

    -- Insert into role_permission join table
    INSERT INTO role_permission (role_id, permission_id)
    VALUES (roleId, permissionId);
END $$;
