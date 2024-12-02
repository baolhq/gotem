package lib

import (
	"fmt"
	"strings"
)

type Config struct {
	Stash Stash `toml:"stash"`
}

type Stash struct {
	Backup   bool               `toml:"backup"`
	Create   bool               `toml:"create"`
	Dotpath  string             `toml:"dotpath"`
	Uselink  bool               `toml:"uselink"`
	Keepdot  bool               `toml:"keepdot"`
	Profiles map[string]Profile `toml:"-"`
}

type Profile struct {
	Backup  *bool           `toml:"backup,omitempty"`
	Create  *bool           `toml:"create,omitempty"`
	Dotpath *string         `toml:"dotpath,omitempty"`
	Uselink *bool           `toml:"uselink,omitempty"`
	Keepdot *bool           `toml:"keepdot,omitempty"`
	Files   map[string]File `toml:"-"`
}

type File struct {
	LocalPath string `toml:"local_path"`
	StashPath string `toml:"stash_path"`
}

func (s *Stash) UnmarshalTOML(data any) error {
	// Ensure `data` is a map
	raw, ok := data.(map[string]interface{})
	if !ok {
		return fmt.Errorf("expected map[string]interface{}, got %T", data)
	}

	// Extract top-level stash settings
	if backup, ok := raw["backup"].(bool); ok {
		s.Backup = backup
	}
	if create, ok := raw["create"].(bool); ok {
		s.Create = create
	}
	if dotpath, ok := raw["dotpath"].(string); ok {
		s.Dotpath = dotpath
	}
	if uselink, ok := raw["uselink"].(bool); ok {
		s.Uselink = uselink
	}
	if keepdot, ok := raw["keepdot"].(bool); ok {
		s.Keepdot = keepdot
	}

	// Parse profiles
	s.Profiles = make(map[string]Profile)
	for key, value := range raw {
		// Skip top-level keys
		if key == "backup" || key == "create" || key == "dotpath" || key == "uselink" || key == "keepdot" {
			continue
		}

		profileData, ok := value.(map[string]interface{})
		if !ok {
			continue
		}

		profile := Profile{
			Files: make(map[string]File),
		}

		// Parse profile-level settings and nested files
		for subKey, subValue := range profileData {
			if subKey == "backup" {
				backup := subValue.(bool)
				profile.Backup = &backup
			} else if subKey == "dotpath" {
				dotpath := subValue.(string)
				profile.Dotpath = &dotpath
			} else {
				// Assume this is a file configuration
				fileData, ok := subValue.(map[string]interface{})
				if !ok {
					continue
				}
				file := File{
					LocalPath: fileData["local_path"].(string),
					StashPath: fileData["stash_path"].(string),
				}
				profile.Files[subKey] = file
			}
		}

		s.Profiles[key] = profile
	}

	return nil
}

func (s Stash) String() string {
	var profilesStr []string
	for key, profile := range s.Profiles {
		profilesStr = append(profilesStr, fmt.Sprintf("Profile: %s\n%s", key, profile.String()))
	}
	return fmt.Sprintf("Backup: %v\nCreate: %v\nDotpath: %s\nUselink: %v\nKeepdot: %v\nProfiles:\n%s",
		s.Backup, s.Create, s.Dotpath, s.Uselink, s.Keepdot, strings.Join(profilesStr, "\n"))
}

// String method for ProfileConfig
func (p Profile) String() string {
	var filesStr []string
	for key, file := range p.Files {
		filesStr = append(filesStr, fmt.Sprintf("  %s: %s", key, file.String()))
	}

	var backupStr, dotpathStr string
	if p.Backup != nil {
		backupStr = fmt.Sprintf("Backup: %v", *p.Backup)
	}
	if p.Dotpath != nil {
		dotpathStr = fmt.Sprintf("Dotpath: %s", *p.Dotpath)
	}

	return fmt.Sprintf("%s\n%s\nFiles:\n%s", backupStr, dotpathStr, strings.Join(filesStr, "\n"))
}

// String method for FileConfig
func (f File) String() string {
	return fmt.Sprintf("LocalPath: %s\nStashPath: %s", f.LocalPath, f.StashPath)
}
