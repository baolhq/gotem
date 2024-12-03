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

func (c Config) String() string {
	return fmt.Sprintf("Stash:\n%v\n", c.Stash)
}

func (s Stash) String() string {
	var profilesStr []string
	for key, profile := range s.Profiles {
		profilesStr = append(profilesStr, fmt.Sprintf("\n    %s\n%s", key, profile.String()))
	}
	return fmt.Sprintf("  Backup: %v\n  Create: %v\n  Dotpath: %s\n  Uselink: %v\n  Keepdot: %v\n  Profiles:%s",
		s.Backup, s.Create, s.Dotpath, s.Uselink, s.Keepdot, strings.Join(profilesStr, ""))
}

func (p Profile) String() string {
	var filesStr []string
	for key, file := range p.Files {
		filesStr = append(filesStr, fmt.Sprintf("  %s: %s", key, file.String()))
	}

	var backupStr, createStr, dotpathStr, keepdotStr, useLinkStr string
	if p.Backup != nil {
		backupStr = fmt.Sprintf("      Backup: %v\n", *p.Backup)
	}
	if p.Create != nil {
		createStr = fmt.Sprintf("      Create: %v\n", *p.Create)
	}
	if p.Dotpath != nil {
		dotpathStr = fmt.Sprintf("      Dotpath: %s\n", *p.Dotpath)
	}
	if p.Keepdot != nil {
		keepdotStr = fmt.Sprintf("      Keep Dots: %v\n", *p.Keepdot)
	}
	if p.Uselink != nil {
		useLinkStr = fmt.Sprintf("      Use link: %v\n", *p.Uselink)
	}

	return fmt.Sprintf("%s%s%s%s%s      Files:\n      %s",
		backupStr, createStr, dotpathStr,
		keepdotStr, useLinkStr, strings.Join(filesStr, "\n      "))
}

func (f File) String() string {
	return fmt.Sprintf("\n          LocalPath: %s\n          StashPath: %s", f.LocalPath, f.StashPath)
}
