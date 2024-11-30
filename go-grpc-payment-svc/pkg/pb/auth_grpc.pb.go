// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: pkg/pb/auth.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	PhoneNumberVerification(ctx context.Context, in *PhoneNumberVerificationRequest, opts ...grpc.CallOption) (*PhoneNumberVerificationResponse, error)
	PhoneNumberVerificationWithOTP(ctx context.Context, in *PhoneNumberVerificationWithOTPRequest, opts ...grpc.CallOption) (*PhoneNumberVerificationWithOTPResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	VerifyResetToken(ctx context.Context, in *VerifyResetTokenRequest, opts ...grpc.CallOption) (*VerifyResetTokenResponse, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error)
	ResendOTP(ctx context.Context, in *ResendOTPRequest, opts ...grpc.CallOption) (*ResendOTPResponse, error)
	UpdatePhoneNumber(ctx context.Context, in *UpdatePhoneNumberRequest, opts ...grpc.CallOption) (*UpdatePhoneNumberResponse, error)
	DeactivateUser(ctx context.Context, in *DeactivateUserRequest, opts ...grpc.CallOption) (*DeactivateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	UpdateDocument(ctx context.Context, in *UpdateDocumentRequest, opts ...grpc.CallOption) (*UpdateDocumentResponse, error)
	KYCVerificationResultCallback(ctx context.Context, in *KYCVerificationResultRequest, opts ...grpc.CallOption) (*KYCVerificationResultResponse, error)
	ValidateEmail(ctx context.Context, in *ValidateEmailRequest, opts ...grpc.CallOption) (*ValidateEmailResponse, error)
	GetUserMetrics(ctx context.Context, in *GetUserMetricsRequest, opts ...grpc.CallOption) (*GetUserMetricsResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) PhoneNumberVerification(ctx context.Context, in *PhoneNumberVerificationRequest, opts ...grpc.CallOption) (*PhoneNumberVerificationResponse, error) {
	out := new(PhoneNumberVerificationResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/PhoneNumberVerification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) PhoneNumberVerificationWithOTP(ctx context.Context, in *PhoneNumberVerificationWithOTPRequest, opts ...grpc.CallOption) (*PhoneNumberVerificationWithOTPResponse, error) {
	out := new(PhoneNumberVerificationWithOTPResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/PhoneNumberVerificationWithOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyResetToken(ctx context.Context, in *VerifyResetTokenRequest, opts ...grpc.CallOption) (*VerifyResetTokenResponse, error) {
	out := new(VerifyResetTokenResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/VerifyResetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error) {
	out := new(UpdatePasswordResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ResendOTP(ctx context.Context, in *ResendOTPRequest, opts ...grpc.CallOption) (*ResendOTPResponse, error) {
	out := new(ResendOTPResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/ResendOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdatePhoneNumber(ctx context.Context, in *UpdatePhoneNumberRequest, opts ...grpc.CallOption) (*UpdatePhoneNumberResponse, error) {
	out := new(UpdatePhoneNumberResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/UpdatePhoneNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeactivateUser(ctx context.Context, in *DeactivateUserRequest, opts ...grpc.CallOption) (*DeactivateUserResponse, error) {
	out := new(DeactivateUserResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/DeactivateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UpdateDocument(ctx context.Context, in *UpdateDocumentRequest, opts ...grpc.CallOption) (*UpdateDocumentResponse, error) {
	out := new(UpdateDocumentResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/UpdateDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) KYCVerificationResultCallback(ctx context.Context, in *KYCVerificationResultRequest, opts ...grpc.CallOption) (*KYCVerificationResultResponse, error) {
	out := new(KYCVerificationResultResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/KYCVerificationResultCallback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ValidateEmail(ctx context.Context, in *ValidateEmailRequest, opts ...grpc.CallOption) (*ValidateEmailResponse, error) {
	out := new(ValidateEmailResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/ValidateEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUserMetrics(ctx context.Context, in *GetUserMetricsRequest, opts ...grpc.CallOption) (*GetUserMetricsResponse, error) {
	out := new(GetUserMetricsResponse)
	err := c.cc.Invoke(ctx, "/auth.AuthService/GetUserMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	PhoneNumberVerification(context.Context, *PhoneNumberVerificationRequest) (*PhoneNumberVerificationResponse, error)
	PhoneNumberVerificationWithOTP(context.Context, *PhoneNumberVerificationWithOTPRequest) (*PhoneNumberVerificationWithOTPResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	VerifyResetToken(context.Context, *VerifyResetTokenRequest) (*VerifyResetTokenResponse, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error)
	ResendOTP(context.Context, *ResendOTPRequest) (*ResendOTPResponse, error)
	UpdatePhoneNumber(context.Context, *UpdatePhoneNumberRequest) (*UpdatePhoneNumberResponse, error)
	DeactivateUser(context.Context, *DeactivateUserRequest) (*DeactivateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	UpdateDocument(context.Context, *UpdateDocumentRequest) (*UpdateDocumentResponse, error)
	KYCVerificationResultCallback(context.Context, *KYCVerificationResultRequest) (*KYCVerificationResultResponse, error)
	ValidateEmail(context.Context, *ValidateEmailRequest) (*ValidateEmailResponse, error)
	GetUserMetrics(context.Context, *GetUserMetricsRequest) (*GetUserMetricsResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) PhoneNumberVerification(context.Context, *PhoneNumberVerificationRequest) (*PhoneNumberVerificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PhoneNumberVerification not implemented")
}
func (UnimplementedAuthServiceServer) PhoneNumberVerificationWithOTP(context.Context, *PhoneNumberVerificationWithOTPRequest) (*PhoneNumberVerificationWithOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PhoneNumberVerificationWithOTP not implemented")
}
func (UnimplementedAuthServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedAuthServiceServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedAuthServiceServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedAuthServiceServer) VerifyResetToken(context.Context, *VerifyResetTokenRequest) (*VerifyResetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyResetToken not implemented")
}
func (UnimplementedAuthServiceServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedAuthServiceServer) ResendOTP(context.Context, *ResendOTPRequest) (*ResendOTPResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResendOTP not implemented")
}
func (UnimplementedAuthServiceServer) UpdatePhoneNumber(context.Context, *UpdatePhoneNumberRequest) (*UpdatePhoneNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePhoneNumber not implemented")
}
func (UnimplementedAuthServiceServer) DeactivateUser(context.Context, *DeactivateUserRequest) (*DeactivateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateUser not implemented")
}
func (UnimplementedAuthServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedAuthServiceServer) UpdateDocument(context.Context, *UpdateDocumentRequest) (*UpdateDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDocument not implemented")
}
func (UnimplementedAuthServiceServer) KYCVerificationResultCallback(context.Context, *KYCVerificationResultRequest) (*KYCVerificationResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KYCVerificationResultCallback not implemented")
}
func (UnimplementedAuthServiceServer) ValidateEmail(context.Context, *ValidateEmailRequest) (*ValidateEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateEmail not implemented")
}
func (UnimplementedAuthServiceServer) GetUserMetrics(context.Context, *GetUserMetricsRequest) (*GetUserMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserMetrics not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_PhoneNumberVerification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneNumberVerificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).PhoneNumberVerification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/PhoneNumberVerification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).PhoneNumberVerification(ctx, req.(*PhoneNumberVerificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_PhoneNumberVerificationWithOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhoneNumberVerificationWithOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).PhoneNumberVerificationWithOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/PhoneNumberVerificationWithOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).PhoneNumberVerificationWithOTP(ctx, req.(*PhoneNumberVerificationWithOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyResetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyResetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyResetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/VerifyResetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyResetToken(ctx, req.(*VerifyResetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ResendOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResendOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ResendOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/ResendOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ResendOTP(ctx, req.(*ResendOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdatePhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePhoneNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdatePhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/UpdatePhoneNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdatePhoneNumber(ctx, req.(*UpdatePhoneNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeactivateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeactivateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/DeactivateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeactivateUser(ctx, req.(*DeactivateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UpdateDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UpdateDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/UpdateDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UpdateDocument(ctx, req.(*UpdateDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_KYCVerificationResultCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KYCVerificationResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).KYCVerificationResultCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/KYCVerificationResultCallback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).KYCVerificationResultCallback(ctx, req.(*KYCVerificationResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ValidateEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ValidateEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/ValidateEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ValidateEmail(ctx, req.(*ValidateEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUserMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthService/GetUserMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserMetrics(ctx, req.(*GetUserMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PhoneNumberVerification",
			Handler:    _AuthService_PhoneNumberVerification_Handler,
		},
		{
			MethodName: "PhoneNumberVerificationWithOTP",
			Handler:    _AuthService_PhoneNumberVerificationWithOTP_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _AuthService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _AuthService_GetUser_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _AuthService_Validate_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _AuthService_ResetPassword_Handler,
		},
		{
			MethodName: "VerifyResetToken",
			Handler:    _AuthService_VerifyResetToken_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _AuthService_UpdatePassword_Handler,
		},
		{
			MethodName: "ResendOTP",
			Handler:    _AuthService_ResendOTP_Handler,
		},
		{
			MethodName: "UpdatePhoneNumber",
			Handler:    _AuthService_UpdatePhoneNumber_Handler,
		},
		{
			MethodName: "DeactivateUser",
			Handler:    _AuthService_DeactivateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _AuthService_DeleteUser_Handler,
		},
		{
			MethodName: "UpdateDocument",
			Handler:    _AuthService_UpdateDocument_Handler,
		},
		{
			MethodName: "KYCVerificationResultCallback",
			Handler:    _AuthService_KYCVerificationResultCallback_Handler,
		},
		{
			MethodName: "ValidateEmail",
			Handler:    _AuthService_ValidateEmail_Handler,
		},
		{
			MethodName: "GetUserMetrics",
			Handler:    _AuthService_GetUserMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/auth.proto",
}
