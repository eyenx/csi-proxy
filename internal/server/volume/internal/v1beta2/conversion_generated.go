// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1beta2

import (
	unsafe "unsafe"

	v1beta2 "github.com/kubernetes-csi/csi-proxy/client/api/volume/v1beta2"
	internal "github.com/kubernetes-csi/csi-proxy/internal/server/volume/internal"
)

func autoConvert_v1beta2_DismountVolumeRequest_To_internal_DismountVolumeRequest(in *v1beta2.DismountVolumeRequest, out *internal.DismountVolumeRequest) error {
	out.VolumeId = in.VolumeId
	out.Path = in.Path
	return nil
}

// Convert_v1beta2_DismountVolumeRequest_To_internal_DismountVolumeRequest is an autogenerated conversion function.
func Convert_v1beta2_DismountVolumeRequest_To_internal_DismountVolumeRequest(in *v1beta2.DismountVolumeRequest, out *internal.DismountVolumeRequest) error {
	return autoConvert_v1beta2_DismountVolumeRequest_To_internal_DismountVolumeRequest(in, out)
}

func autoConvert_internal_DismountVolumeRequest_To_v1beta2_DismountVolumeRequest(in *internal.DismountVolumeRequest, out *v1beta2.DismountVolumeRequest) error {
	out.VolumeId = in.VolumeId
	out.Path = in.Path
	return nil
}

// Convert_internal_DismountVolumeRequest_To_v1beta2_DismountVolumeRequest is an autogenerated conversion function.
func Convert_internal_DismountVolumeRequest_To_v1beta2_DismountVolumeRequest(in *internal.DismountVolumeRequest, out *v1beta2.DismountVolumeRequest) error {
	return autoConvert_internal_DismountVolumeRequest_To_v1beta2_DismountVolumeRequest(in, out)
}

func autoConvert_v1beta2_DismountVolumeResponse_To_internal_DismountVolumeResponse(in *v1beta2.DismountVolumeResponse, out *internal.DismountVolumeResponse) error {
	return nil
}

// Convert_v1beta2_DismountVolumeResponse_To_internal_DismountVolumeResponse is an autogenerated conversion function.
func Convert_v1beta2_DismountVolumeResponse_To_internal_DismountVolumeResponse(in *v1beta2.DismountVolumeResponse, out *internal.DismountVolumeResponse) error {
	return autoConvert_v1beta2_DismountVolumeResponse_To_internal_DismountVolumeResponse(in, out)
}

func autoConvert_internal_DismountVolumeResponse_To_v1beta2_DismountVolumeResponse(in *internal.DismountVolumeResponse, out *v1beta2.DismountVolumeResponse) error {
	return nil
}

// Convert_internal_DismountVolumeResponse_To_v1beta2_DismountVolumeResponse is an autogenerated conversion function.
func Convert_internal_DismountVolumeResponse_To_v1beta2_DismountVolumeResponse(in *internal.DismountVolumeResponse, out *v1beta2.DismountVolumeResponse) error {
	return autoConvert_internal_DismountVolumeResponse_To_v1beta2_DismountVolumeResponse(in, out)
}

func autoConvert_v1beta2_FormatVolumeRequest_To_internal_FormatVolumeRequest(in *v1beta2.FormatVolumeRequest, out *internal.FormatVolumeRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_v1beta2_FormatVolumeRequest_To_internal_FormatVolumeRequest is an autogenerated conversion function.
func Convert_v1beta2_FormatVolumeRequest_To_internal_FormatVolumeRequest(in *v1beta2.FormatVolumeRequest, out *internal.FormatVolumeRequest) error {
	return autoConvert_v1beta2_FormatVolumeRequest_To_internal_FormatVolumeRequest(in, out)
}

func autoConvert_internal_FormatVolumeRequest_To_v1beta2_FormatVolumeRequest(in *internal.FormatVolumeRequest, out *v1beta2.FormatVolumeRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_internal_FormatVolumeRequest_To_v1beta2_FormatVolumeRequest is an autogenerated conversion function.
func Convert_internal_FormatVolumeRequest_To_v1beta2_FormatVolumeRequest(in *internal.FormatVolumeRequest, out *v1beta2.FormatVolumeRequest) error {
	return autoConvert_internal_FormatVolumeRequest_To_v1beta2_FormatVolumeRequest(in, out)
}

func autoConvert_v1beta2_FormatVolumeResponse_To_internal_FormatVolumeResponse(in *v1beta2.FormatVolumeResponse, out *internal.FormatVolumeResponse) error {
	return nil
}

// Convert_v1beta2_FormatVolumeResponse_To_internal_FormatVolumeResponse is an autogenerated conversion function.
func Convert_v1beta2_FormatVolumeResponse_To_internal_FormatVolumeResponse(in *v1beta2.FormatVolumeResponse, out *internal.FormatVolumeResponse) error {
	return autoConvert_v1beta2_FormatVolumeResponse_To_internal_FormatVolumeResponse(in, out)
}

func autoConvert_internal_FormatVolumeResponse_To_v1beta2_FormatVolumeResponse(in *internal.FormatVolumeResponse, out *v1beta2.FormatVolumeResponse) error {
	return nil
}

// Convert_internal_FormatVolumeResponse_To_v1beta2_FormatVolumeResponse is an autogenerated conversion function.
func Convert_internal_FormatVolumeResponse_To_v1beta2_FormatVolumeResponse(in *internal.FormatVolumeResponse, out *v1beta2.FormatVolumeResponse) error {
	return autoConvert_internal_FormatVolumeResponse_To_v1beta2_FormatVolumeResponse(in, out)
}

func autoConvert_v1beta2_IsVolumeFormattedRequest_To_internal_IsVolumeFormattedRequest(in *v1beta2.IsVolumeFormattedRequest, out *internal.IsVolumeFormattedRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_v1beta2_IsVolumeFormattedRequest_To_internal_IsVolumeFormattedRequest is an autogenerated conversion function.
func Convert_v1beta2_IsVolumeFormattedRequest_To_internal_IsVolumeFormattedRequest(in *v1beta2.IsVolumeFormattedRequest, out *internal.IsVolumeFormattedRequest) error {
	return autoConvert_v1beta2_IsVolumeFormattedRequest_To_internal_IsVolumeFormattedRequest(in, out)
}

func autoConvert_internal_IsVolumeFormattedRequest_To_v1beta2_IsVolumeFormattedRequest(in *internal.IsVolumeFormattedRequest, out *v1beta2.IsVolumeFormattedRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_internal_IsVolumeFormattedRequest_To_v1beta2_IsVolumeFormattedRequest is an autogenerated conversion function.
func Convert_internal_IsVolumeFormattedRequest_To_v1beta2_IsVolumeFormattedRequest(in *internal.IsVolumeFormattedRequest, out *v1beta2.IsVolumeFormattedRequest) error {
	return autoConvert_internal_IsVolumeFormattedRequest_To_v1beta2_IsVolumeFormattedRequest(in, out)
}

func autoConvert_v1beta2_IsVolumeFormattedResponse_To_internal_IsVolumeFormattedResponse(in *v1beta2.IsVolumeFormattedResponse, out *internal.IsVolumeFormattedResponse) error {
	out.Formatted = in.Formatted
	return nil
}

// Convert_v1beta2_IsVolumeFormattedResponse_To_internal_IsVolumeFormattedResponse is an autogenerated conversion function.
func Convert_v1beta2_IsVolumeFormattedResponse_To_internal_IsVolumeFormattedResponse(in *v1beta2.IsVolumeFormattedResponse, out *internal.IsVolumeFormattedResponse) error {
	return autoConvert_v1beta2_IsVolumeFormattedResponse_To_internal_IsVolumeFormattedResponse(in, out)
}

func autoConvert_internal_IsVolumeFormattedResponse_To_v1beta2_IsVolumeFormattedResponse(in *internal.IsVolumeFormattedResponse, out *v1beta2.IsVolumeFormattedResponse) error {
	out.Formatted = in.Formatted
	return nil
}

// Convert_internal_IsVolumeFormattedResponse_To_v1beta2_IsVolumeFormattedResponse is an autogenerated conversion function.
func Convert_internal_IsVolumeFormattedResponse_To_v1beta2_IsVolumeFormattedResponse(in *internal.IsVolumeFormattedResponse, out *v1beta2.IsVolumeFormattedResponse) error {
	return autoConvert_internal_IsVolumeFormattedResponse_To_v1beta2_IsVolumeFormattedResponse(in, out)
}

func autoConvert_v1beta2_ListVolumesOnDiskRequest_To_internal_ListVolumesOnDiskRequest(in *v1beta2.ListVolumesOnDiskRequest, out *internal.ListVolumesOnDiskRequest) error {
	out.DiskId = in.DiskId
	return nil
}

// Convert_v1beta2_ListVolumesOnDiskRequest_To_internal_ListVolumesOnDiskRequest is an autogenerated conversion function.
func Convert_v1beta2_ListVolumesOnDiskRequest_To_internal_ListVolumesOnDiskRequest(in *v1beta2.ListVolumesOnDiskRequest, out *internal.ListVolumesOnDiskRequest) error {
	return autoConvert_v1beta2_ListVolumesOnDiskRequest_To_internal_ListVolumesOnDiskRequest(in, out)
}

func autoConvert_internal_ListVolumesOnDiskRequest_To_v1beta2_ListVolumesOnDiskRequest(in *internal.ListVolumesOnDiskRequest, out *v1beta2.ListVolumesOnDiskRequest) error {
	out.DiskId = in.DiskId
	return nil
}

// Convert_internal_ListVolumesOnDiskRequest_To_v1beta2_ListVolumesOnDiskRequest is an autogenerated conversion function.
func Convert_internal_ListVolumesOnDiskRequest_To_v1beta2_ListVolumesOnDiskRequest(in *internal.ListVolumesOnDiskRequest, out *v1beta2.ListVolumesOnDiskRequest) error {
	return autoConvert_internal_ListVolumesOnDiskRequest_To_v1beta2_ListVolumesOnDiskRequest(in, out)
}

func autoConvert_v1beta2_ListVolumesOnDiskResponse_To_internal_ListVolumesOnDiskResponse(in *v1beta2.ListVolumesOnDiskResponse, out *internal.ListVolumesOnDiskResponse) error {
	out.VolumeIds = *(*[]string)(unsafe.Pointer(&in.VolumeIds))
	return nil
}

// Convert_v1beta2_ListVolumesOnDiskResponse_To_internal_ListVolumesOnDiskResponse is an autogenerated conversion function.
func Convert_v1beta2_ListVolumesOnDiskResponse_To_internal_ListVolumesOnDiskResponse(in *v1beta2.ListVolumesOnDiskResponse, out *internal.ListVolumesOnDiskResponse) error {
	return autoConvert_v1beta2_ListVolumesOnDiskResponse_To_internal_ListVolumesOnDiskResponse(in, out)
}

func autoConvert_internal_ListVolumesOnDiskResponse_To_v1beta2_ListVolumesOnDiskResponse(in *internal.ListVolumesOnDiskResponse, out *v1beta2.ListVolumesOnDiskResponse) error {
	out.VolumeIds = *(*[]string)(unsafe.Pointer(&in.VolumeIds))
	return nil
}

// Convert_internal_ListVolumesOnDiskResponse_To_v1beta2_ListVolumesOnDiskResponse is an autogenerated conversion function.
func Convert_internal_ListVolumesOnDiskResponse_To_v1beta2_ListVolumesOnDiskResponse(in *internal.ListVolumesOnDiskResponse, out *v1beta2.ListVolumesOnDiskResponse) error {
	return autoConvert_internal_ListVolumesOnDiskResponse_To_v1beta2_ListVolumesOnDiskResponse(in, out)
}

func autoConvert_v1beta2_MountVolumeRequest_To_internal_MountVolumeRequest(in *v1beta2.MountVolumeRequest, out *internal.MountVolumeRequest) error {
	out.VolumeId = in.VolumeId
	out.Path = in.Path
	return nil
}

// Convert_v1beta2_MountVolumeRequest_To_internal_MountVolumeRequest is an autogenerated conversion function.
func Convert_v1beta2_MountVolumeRequest_To_internal_MountVolumeRequest(in *v1beta2.MountVolumeRequest, out *internal.MountVolumeRequest) error {
	return autoConvert_v1beta2_MountVolumeRequest_To_internal_MountVolumeRequest(in, out)
}

func autoConvert_internal_MountVolumeRequest_To_v1beta2_MountVolumeRequest(in *internal.MountVolumeRequest, out *v1beta2.MountVolumeRequest) error {
	out.VolumeId = in.VolumeId
	out.Path = in.Path
	return nil
}

// Convert_internal_MountVolumeRequest_To_v1beta2_MountVolumeRequest is an autogenerated conversion function.
func Convert_internal_MountVolumeRequest_To_v1beta2_MountVolumeRequest(in *internal.MountVolumeRequest, out *v1beta2.MountVolumeRequest) error {
	return autoConvert_internal_MountVolumeRequest_To_v1beta2_MountVolumeRequest(in, out)
}

func autoConvert_v1beta2_MountVolumeResponse_To_internal_MountVolumeResponse(in *v1beta2.MountVolumeResponse, out *internal.MountVolumeResponse) error {
	return nil
}

// Convert_v1beta2_MountVolumeResponse_To_internal_MountVolumeResponse is an autogenerated conversion function.
func Convert_v1beta2_MountVolumeResponse_To_internal_MountVolumeResponse(in *v1beta2.MountVolumeResponse, out *internal.MountVolumeResponse) error {
	return autoConvert_v1beta2_MountVolumeResponse_To_internal_MountVolumeResponse(in, out)
}

func autoConvert_internal_MountVolumeResponse_To_v1beta2_MountVolumeResponse(in *internal.MountVolumeResponse, out *v1beta2.MountVolumeResponse) error {
	return nil
}

// Convert_internal_MountVolumeResponse_To_v1beta2_MountVolumeResponse is an autogenerated conversion function.
func Convert_internal_MountVolumeResponse_To_v1beta2_MountVolumeResponse(in *internal.MountVolumeResponse, out *v1beta2.MountVolumeResponse) error {
	return autoConvert_internal_MountVolumeResponse_To_v1beta2_MountVolumeResponse(in, out)
}

func autoConvert_v1beta2_ResizeVolumeRequest_To_internal_ResizeVolumeRequest(in *v1beta2.ResizeVolumeRequest, out *internal.ResizeVolumeRequest) error {
	out.VolumeId = in.VolumeId
	out.Size = in.Size
	return nil
}

// Convert_v1beta2_ResizeVolumeRequest_To_internal_ResizeVolumeRequest is an autogenerated conversion function.
func Convert_v1beta2_ResizeVolumeRequest_To_internal_ResizeVolumeRequest(in *v1beta2.ResizeVolumeRequest, out *internal.ResizeVolumeRequest) error {
	return autoConvert_v1beta2_ResizeVolumeRequest_To_internal_ResizeVolumeRequest(in, out)
}

func autoConvert_internal_ResizeVolumeRequest_To_v1beta2_ResizeVolumeRequest(in *internal.ResizeVolumeRequest, out *v1beta2.ResizeVolumeRequest) error {
	out.VolumeId = in.VolumeId
	out.Size = in.Size
	return nil
}

// Convert_internal_ResizeVolumeRequest_To_v1beta2_ResizeVolumeRequest is an autogenerated conversion function.
func Convert_internal_ResizeVolumeRequest_To_v1beta2_ResizeVolumeRequest(in *internal.ResizeVolumeRequest, out *v1beta2.ResizeVolumeRequest) error {
	return autoConvert_internal_ResizeVolumeRequest_To_v1beta2_ResizeVolumeRequest(in, out)
}

func autoConvert_v1beta2_ResizeVolumeResponse_To_internal_ResizeVolumeResponse(in *v1beta2.ResizeVolumeResponse, out *internal.ResizeVolumeResponse) error {
	return nil
}

// Convert_v1beta2_ResizeVolumeResponse_To_internal_ResizeVolumeResponse is an autogenerated conversion function.
func Convert_v1beta2_ResizeVolumeResponse_To_internal_ResizeVolumeResponse(in *v1beta2.ResizeVolumeResponse, out *internal.ResizeVolumeResponse) error {
	return autoConvert_v1beta2_ResizeVolumeResponse_To_internal_ResizeVolumeResponse(in, out)
}

func autoConvert_internal_ResizeVolumeResponse_To_v1beta2_ResizeVolumeResponse(in *internal.ResizeVolumeResponse, out *v1beta2.ResizeVolumeResponse) error {
	return nil
}

// Convert_internal_ResizeVolumeResponse_To_v1beta2_ResizeVolumeResponse is an autogenerated conversion function.
func Convert_internal_ResizeVolumeResponse_To_v1beta2_ResizeVolumeResponse(in *internal.ResizeVolumeResponse, out *v1beta2.ResizeVolumeResponse) error {
	return autoConvert_internal_ResizeVolumeResponse_To_v1beta2_ResizeVolumeResponse(in, out)
}

func autoConvert_v1beta2_VolumeDiskNumberRequest_To_internal_VolumeDiskNumberRequest(in *v1beta2.VolumeDiskNumberRequest, out *internal.VolumeDiskNumberRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_v1beta2_VolumeDiskNumberRequest_To_internal_VolumeDiskNumberRequest is an autogenerated conversion function.
func Convert_v1beta2_VolumeDiskNumberRequest_To_internal_VolumeDiskNumberRequest(in *v1beta2.VolumeDiskNumberRequest, out *internal.VolumeDiskNumberRequest) error {
	return autoConvert_v1beta2_VolumeDiskNumberRequest_To_internal_VolumeDiskNumberRequest(in, out)
}

func autoConvert_internal_VolumeDiskNumberRequest_To_v1beta2_VolumeDiskNumberRequest(in *internal.VolumeDiskNumberRequest, out *v1beta2.VolumeDiskNumberRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_internal_VolumeDiskNumberRequest_To_v1beta2_VolumeDiskNumberRequest is an autogenerated conversion function.
func Convert_internal_VolumeDiskNumberRequest_To_v1beta2_VolumeDiskNumberRequest(in *internal.VolumeDiskNumberRequest, out *v1beta2.VolumeDiskNumberRequest) error {
	return autoConvert_internal_VolumeDiskNumberRequest_To_v1beta2_VolumeDiskNumberRequest(in, out)
}

func autoConvert_v1beta2_VolumeDiskNumberResponse_To_internal_VolumeDiskNumberResponse(in *v1beta2.VolumeDiskNumberResponse, out *internal.VolumeDiskNumberResponse) error {
	out.DiskNumber = in.DiskNumber
	return nil
}

// Convert_v1beta2_VolumeDiskNumberResponse_To_internal_VolumeDiskNumberResponse is an autogenerated conversion function.
func Convert_v1beta2_VolumeDiskNumberResponse_To_internal_VolumeDiskNumberResponse(in *v1beta2.VolumeDiskNumberResponse, out *internal.VolumeDiskNumberResponse) error {
	return autoConvert_v1beta2_VolumeDiskNumberResponse_To_internal_VolumeDiskNumberResponse(in, out)
}

func autoConvert_internal_VolumeDiskNumberResponse_To_v1beta2_VolumeDiskNumberResponse(in *internal.VolumeDiskNumberResponse, out *v1beta2.VolumeDiskNumberResponse) error {
	out.DiskNumber = in.DiskNumber
	return nil
}

// Convert_internal_VolumeDiskNumberResponse_To_v1beta2_VolumeDiskNumberResponse is an autogenerated conversion function.
func Convert_internal_VolumeDiskNumberResponse_To_v1beta2_VolumeDiskNumberResponse(in *internal.VolumeDiskNumberResponse, out *v1beta2.VolumeDiskNumberResponse) error {
	return autoConvert_internal_VolumeDiskNumberResponse_To_v1beta2_VolumeDiskNumberResponse(in, out)
}

func autoConvert_v1beta2_VolumeIDFromMountRequest_To_internal_VolumeIDFromMountRequest(in *v1beta2.VolumeIDFromMountRequest, out *internal.VolumeIDFromMountRequest) error {
	out.Mount = in.Mount
	return nil
}

// Convert_v1beta2_VolumeIDFromMountRequest_To_internal_VolumeIDFromMountRequest is an autogenerated conversion function.
func Convert_v1beta2_VolumeIDFromMountRequest_To_internal_VolumeIDFromMountRequest(in *v1beta2.VolumeIDFromMountRequest, out *internal.VolumeIDFromMountRequest) error {
	return autoConvert_v1beta2_VolumeIDFromMountRequest_To_internal_VolumeIDFromMountRequest(in, out)
}

func autoConvert_internal_VolumeIDFromMountRequest_To_v1beta2_VolumeIDFromMountRequest(in *internal.VolumeIDFromMountRequest, out *v1beta2.VolumeIDFromMountRequest) error {
	out.Mount = in.Mount
	return nil
}

// Convert_internal_VolumeIDFromMountRequest_To_v1beta2_VolumeIDFromMountRequest is an autogenerated conversion function.
func Convert_internal_VolumeIDFromMountRequest_To_v1beta2_VolumeIDFromMountRequest(in *internal.VolumeIDFromMountRequest, out *v1beta2.VolumeIDFromMountRequest) error {
	return autoConvert_internal_VolumeIDFromMountRequest_To_v1beta2_VolumeIDFromMountRequest(in, out)
}

func autoConvert_v1beta2_VolumeIDFromMountResponse_To_internal_VolumeIDFromMountResponse(in *v1beta2.VolumeIDFromMountResponse, out *internal.VolumeIDFromMountResponse) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_v1beta2_VolumeIDFromMountResponse_To_internal_VolumeIDFromMountResponse is an autogenerated conversion function.
func Convert_v1beta2_VolumeIDFromMountResponse_To_internal_VolumeIDFromMountResponse(in *v1beta2.VolumeIDFromMountResponse, out *internal.VolumeIDFromMountResponse) error {
	return autoConvert_v1beta2_VolumeIDFromMountResponse_To_internal_VolumeIDFromMountResponse(in, out)
}

func autoConvert_internal_VolumeIDFromMountResponse_To_v1beta2_VolumeIDFromMountResponse(in *internal.VolumeIDFromMountResponse, out *v1beta2.VolumeIDFromMountResponse) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_internal_VolumeIDFromMountResponse_To_v1beta2_VolumeIDFromMountResponse is an autogenerated conversion function.
func Convert_internal_VolumeIDFromMountResponse_To_v1beta2_VolumeIDFromMountResponse(in *internal.VolumeIDFromMountResponse, out *v1beta2.VolumeIDFromMountResponse) error {
	return autoConvert_internal_VolumeIDFromMountResponse_To_v1beta2_VolumeIDFromMountResponse(in, out)
}

func autoConvert_v1beta2_VolumeStatsRequest_To_internal_VolumeStatsRequest(in *v1beta2.VolumeStatsRequest, out *internal.VolumeStatsRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_v1beta2_VolumeStatsRequest_To_internal_VolumeStatsRequest is an autogenerated conversion function.
func Convert_v1beta2_VolumeStatsRequest_To_internal_VolumeStatsRequest(in *v1beta2.VolumeStatsRequest, out *internal.VolumeStatsRequest) error {
	return autoConvert_v1beta2_VolumeStatsRequest_To_internal_VolumeStatsRequest(in, out)
}

func autoConvert_internal_VolumeStatsRequest_To_v1beta2_VolumeStatsRequest(in *internal.VolumeStatsRequest, out *v1beta2.VolumeStatsRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_internal_VolumeStatsRequest_To_v1beta2_VolumeStatsRequest is an autogenerated conversion function.
func Convert_internal_VolumeStatsRequest_To_v1beta2_VolumeStatsRequest(in *internal.VolumeStatsRequest, out *v1beta2.VolumeStatsRequest) error {
	return autoConvert_internal_VolumeStatsRequest_To_v1beta2_VolumeStatsRequest(in, out)
}

func autoConvert_v1beta2_VolumeStatsResponse_To_internal_VolumeStatsResponse(in *v1beta2.VolumeStatsResponse, out *internal.VolumeStatsResponse) error {
	out.VolumeSize = in.VolumeSize
	out.VolumeUsedSize = in.VolumeUsedSize
	return nil
}

// Convert_v1beta2_VolumeStatsResponse_To_internal_VolumeStatsResponse is an autogenerated conversion function.
func Convert_v1beta2_VolumeStatsResponse_To_internal_VolumeStatsResponse(in *v1beta2.VolumeStatsResponse, out *internal.VolumeStatsResponse) error {
	return autoConvert_v1beta2_VolumeStatsResponse_To_internal_VolumeStatsResponse(in, out)
}

func autoConvert_internal_VolumeStatsResponse_To_v1beta2_VolumeStatsResponse(in *internal.VolumeStatsResponse, out *v1beta2.VolumeStatsResponse) error {
	out.VolumeSize = in.VolumeSize
	out.VolumeUsedSize = in.VolumeUsedSize
	return nil
}

// Convert_internal_VolumeStatsResponse_To_v1beta2_VolumeStatsResponse is an autogenerated conversion function.
func Convert_internal_VolumeStatsResponse_To_v1beta2_VolumeStatsResponse(in *internal.VolumeStatsResponse, out *v1beta2.VolumeStatsResponse) error {
	return autoConvert_internal_VolumeStatsResponse_To_v1beta2_VolumeStatsResponse(in, out)
}

func autoConvert_v1beta2_WriteVolumeCacheRequest_To_internal_WriteVolumeCacheRequest(in *v1beta2.WriteVolumeCacheRequest, out *internal.WriteVolumeCacheRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_v1beta2_WriteVolumeCacheRequest_To_internal_WriteVolumeCacheRequest is an autogenerated conversion function.
func Convert_v1beta2_WriteVolumeCacheRequest_To_internal_WriteVolumeCacheRequest(in *v1beta2.WriteVolumeCacheRequest, out *internal.WriteVolumeCacheRequest) error {
	return autoConvert_v1beta2_WriteVolumeCacheRequest_To_internal_WriteVolumeCacheRequest(in, out)
}

func autoConvert_internal_WriteVolumeCacheRequest_To_v1beta2_WriteVolumeCacheRequest(in *internal.WriteVolumeCacheRequest, out *v1beta2.WriteVolumeCacheRequest) error {
	out.VolumeId = in.VolumeId
	return nil
}

// Convert_internal_WriteVolumeCacheRequest_To_v1beta2_WriteVolumeCacheRequest is an autogenerated conversion function.
func Convert_internal_WriteVolumeCacheRequest_To_v1beta2_WriteVolumeCacheRequest(in *internal.WriteVolumeCacheRequest, out *v1beta2.WriteVolumeCacheRequest) error {
	return autoConvert_internal_WriteVolumeCacheRequest_To_v1beta2_WriteVolumeCacheRequest(in, out)
}

func autoConvert_v1beta2_WriteVolumeCacheResponse_To_internal_WriteVolumeCacheResponse(in *v1beta2.WriteVolumeCacheResponse, out *internal.WriteVolumeCacheResponse) error {
	return nil
}

// Convert_v1beta2_WriteVolumeCacheResponse_To_internal_WriteVolumeCacheResponse is an autogenerated conversion function.
func Convert_v1beta2_WriteVolumeCacheResponse_To_internal_WriteVolumeCacheResponse(in *v1beta2.WriteVolumeCacheResponse, out *internal.WriteVolumeCacheResponse) error {
	return autoConvert_v1beta2_WriteVolumeCacheResponse_To_internal_WriteVolumeCacheResponse(in, out)
}

func autoConvert_internal_WriteVolumeCacheResponse_To_v1beta2_WriteVolumeCacheResponse(in *internal.WriteVolumeCacheResponse, out *v1beta2.WriteVolumeCacheResponse) error {
	return nil
}

// Convert_internal_WriteVolumeCacheResponse_To_v1beta2_WriteVolumeCacheResponse is an autogenerated conversion function.
func Convert_internal_WriteVolumeCacheResponse_To_v1beta2_WriteVolumeCacheResponse(in *internal.WriteVolumeCacheResponse, out *v1beta2.WriteVolumeCacheResponse) error {
	return autoConvert_internal_WriteVolumeCacheResponse_To_v1beta2_WriteVolumeCacheResponse(in, out)
}
