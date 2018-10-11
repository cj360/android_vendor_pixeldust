package android
type Product_variables struct {

	Bootloader_message_offset struct {
		Cflags []string
	}

	Target_shim_libs struct {
		Cppflags []string
	}

	Target_uses_color_metadata struct {
		Cppflags []string
	}
}

type ProductVariables struct {
	Target_shim_libs  *string `json:",omitempty"`
	Target_uses_color_metadata  *bool `json:",omitempty"`
	Bootloader_message_offset  *int `json:",omitempty"`
}
