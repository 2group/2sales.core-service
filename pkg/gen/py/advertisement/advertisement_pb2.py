# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: advertisement/advertisement.proto
# Protobuf Python Version: 5.29.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    0,
    '',
    'advertisement/advertisement.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!advertisement/advertisement.proto\x12\x03\x63rm\"\xe5\x01\n\x0b\x42\x61nnerModel\x12\x11\n\tbanner_id\x18\x01 \x01(\x03\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x18\n\x0b\x64\x65scription\x18\x03 \x01(\tH\x00\x88\x01\x01\x12\x11\n\timage_url\x18\x04 \x01(\t\x12\x17\n\ntarget_url\x18\x05 \x01(\tH\x01\x88\x01\x01\x12\x11\n\tis_active\x18\x06 \x01(\x08\x12\x15\n\rdisplay_order\x18\x07 \x01(\x05\x12\x12\n\ncreated_at\x18\x08 \x01(\t\x12\x12\n\nupdated_at\x18\t \x01(\tB\x0e\n\x0c_descriptionB\r\n\x0b_target_url\"7\n\x13\x43reateBannerRequest\x12 \n\x06\x62\x61nner\x18\x01 \x01(\x0b\x32\x10.crm.BannerModel\"8\n\x14\x43reateBannerResponse\x12 \n\x06\x62\x61nner\x18\x01 \x01(\x0b\x32\x10.crm.BannerModel\"\'\n\x12ListBannersRequest\x12\x11\n\tis_active\x18\x01 \x01(\x08\"8\n\x13ListBannersResponse\x12!\n\x07\x62\x61nners\x18\x01 \x03(\x0b\x32\x10.crm.BannerModel2\x9d\x01\n\x14\x41\x64vertisementService\x12\x43\n\x0c\x43reateBanner\x12\x18.crm.CreateBannerRequest\x1a\x19.crm.CreateBannerResponse\x12@\n\x0bListBanners\x12\x17.crm.ListBannersRequest\x1a\x18.crm.ListBannersResponseBPZNgithub.com/2group/2sales.core-service/pkg/gen/go/advertisement;advertisementv1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'advertisement.advertisement_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZNgithub.com/2group/2sales.core-service/pkg/gen/go/advertisement;advertisementv1'
  _globals['_BANNERMODEL']._serialized_start=43
  _globals['_BANNERMODEL']._serialized_end=272
  _globals['_CREATEBANNERREQUEST']._serialized_start=274
  _globals['_CREATEBANNERREQUEST']._serialized_end=329
  _globals['_CREATEBANNERRESPONSE']._serialized_start=331
  _globals['_CREATEBANNERRESPONSE']._serialized_end=387
  _globals['_LISTBANNERSREQUEST']._serialized_start=389
  _globals['_LISTBANNERSREQUEST']._serialized_end=428
  _globals['_LISTBANNERSRESPONSE']._serialized_start=430
  _globals['_LISTBANNERSRESPONSE']._serialized_end=486
  _globals['_ADVERTISEMENTSERVICE']._serialized_start=489
  _globals['_ADVERTISEMENTSERVICE']._serialized_end=646
# @@protoc_insertion_point(module_scope)
