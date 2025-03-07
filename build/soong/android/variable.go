package android
type Product_variables struct {
	Additional_gralloc_10_usage_bits struct {
		Cppflags []string
	}
	Device_support_hwfde struct {
		Cflags []string
		Header_libs []string
		Shared_libs []string
	}
	Device_support_hwfde_perf struct {
		Cflags []string
	}
	Device_support_legacy_hwfde struct {
		Cflags []string
	}
	Device_support_wait_for_qsee struct {
		Cflags []string
	}
	Has_legacy_camera_hal1 struct {
		Cflags []string
	}

	Needs_text_relocations struct {
		Cppflags []string
	}

	Target_process_sdk_version_override struct {
		Cppflags []string
	}

	Target_shim_libs struct {
		Cppflags []string
	}

	Uses_generic_camera_parameter_library struct {
		Srcs []string
	}

	Uses_qcom_bsp_legacy struct {
		Cppflags []string
	}

	Uses_qti_camera_device struct {
		Cppflags []string
		Shared_libs []string
	}
	Apply_msm8974_1440p_egl_workaround struct {
		Cflags []string
	}
}

type ProductVariables struct {
	Additional_gralloc_10_usage_bits  *string `json:",omitempty"`
	Device_support_hwfde  *bool `json:",omitempty"`
	Device_support_hwfde_perf  *bool `json:",omitempty"`
	Device_support_legacy_hwfde  *bool `json:",omitempty"`
	Device_support_wait_for_qsee  *bool `json:",omitempty"`
	Has_legacy_camera_hal1  *bool `json:",omitempty"`
	Needs_text_relocations  *bool `json:",omitempty"`
	Specific_camera_parameter_library  *string `json:",omitempty"`
	Target_process_sdk_version_override *string `json:",omitempty"`
	Target_shim_libs  *string `json:",omitempty"`
	Uses_generic_camera_parameter_library  *bool `json:",omitempty"`
	Uses_qcom_bsp_legacy  *bool `json:",omitempty"`
	Uses_qti_camera_device  *bool `json:",omitempty"`
	Apply_msm8974_1440p_egl_workaround  *bool `json:",omitempty"`
}
