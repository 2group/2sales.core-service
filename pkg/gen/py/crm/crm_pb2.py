# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: crm/crm.proto
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
    'crm/crm.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from user import user_pb2 as user_dot_user__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rcrm/crm.proto\x12\x03\x63rm\x1a\x0fuser/user.proto\"\xa8\x01\n\x0cOrganization\x12\x1c\n\x0forganization_id\x18\x01 \x01(\x03H\x00\x88\x01\x01\x12\x11\n\x04name\x18\x02 \x01(\tH\x01\x88\x01\x01\x12\x17\n\nlegal_name\x18\x03 \x01(\tH\x02\x88\x01\x01\x12\"\n\x07\x61\x64\x64ress\x18\x04 \x01(\x0b\x32\x11.crm.AddressModelB\x12\n\x10_organization_idB\x07\n\x05_nameB\r\n\x0b_legal_name\"\xde\x03\n\x04Lead\x12\x14\n\x07lead_id\x18\x01 \x01(\x03H\x00\x88\x01\x01\x12\x16\n\tlead_name\x18\x02 \x01(\tH\x01\x88\x01\x01\x12\'\n\x1a\x63reated_by_organization_id\x18\x03 \x01(\x03H\x02\x88\x01\x01\x12\x39\n\x19\x63ounterparty_organization\x18\x04 \x01(\x0b\x32\x11.crm.OrganizationH\x03\x88\x01\x01\x12\x1e\n\x08\x63ontacts\x18\x05 \x03(\x0b\x32\x0c.crm.Contact\x12\r\n\x05stage\x18\x06 \x01(\t\x12\x1c\n\x0f\x65stimated_value\x18\x07 \x01(\x01H\x04\x88\x01\x01\x12(\n\x0f\x63reated_by_user\x18\x08 \x01(\x0b\x32\n.user.UserH\x05\x88\x01\x01\x12\x17\n\ncreated_at\x18\t \x01(\tH\x06\x88\x01\x01\x12\x17\n\nupdated_at\x18\n \x01(\tH\x07\x88\x01\x01\x42\n\n\x08_lead_idB\x0c\n\n_lead_nameB\x1d\n\x1b_created_by_organization_idB\x1c\n\x1a_counterparty_organizationB\x12\n\x10_estimated_valueB\x12\n\x10_created_by_userB\r\n\x0b_created_atB\r\n\x0b_updated_at\"\xff\x01\n\x07\x43ontact\x12\x17\n\ncontact_id\x18\x01 \x01(\x03H\x00\x88\x01\x01\x12\x1c\n\x0forganization_id\x18\x02 \x01(\x03H\x01\x88\x01\x01\x12\x14\n\x07lead_id\x18\x03 \x01(\x03H\x02\x88\x01\x01\x12\x1b\n\x0e\x63ontact_person\x18\x04 \x01(\tH\x03\x88\x01\x01\x12\x19\n\x0cphone_number\x18\x05 \x01(\tH\x04\x88\x01\x01\x12\x12\n\x05\x65mail\x18\x06 \x01(\tH\x05\x88\x01\x01\x42\r\n\x0b_contact_idB\x12\n\x10_organization_idB\n\n\x08_lead_idB\x11\n\x0f_contact_personB\x0f\n\r_phone_numberB\x08\n\x06_email\"\xb0\x03\n\tLeadModel\x12\x14\n\x07lead_id\x18\x01 \x01(\x03H\x00\x88\x01\x01\x12\x16\n\tlead_name\x18\x02 \x01(\tH\x01\x88\x01\x01\x12\'\n\x1a\x63reated_by_organization_id\x18\x03 \x01(\x03H\x02\x88\x01\x01\x12)\n\x1c\x63ounterparty_organization_id\x18\x04 \x01(\x03H\x03\x88\x01\x01\x12\r\n\x05stage\x18\x05 \x01(\t\x12\x1c\n\x0f\x65stimated_value\x18\x06 \x01(\x01H\x04\x88\x01\x01\x12\x1f\n\x12\x63reated_by_user_id\x18\x07 \x01(\x03H\x05\x88\x01\x01\x12\x17\n\ncreated_at\x18\x08 \x01(\tH\x06\x88\x01\x01\x12\x17\n\nupdated_at\x18\t \x01(\tH\x07\x88\x01\x01\x42\n\n\x08_lead_idB\x0c\n\n_lead_nameB\x1d\n\x1b_created_by_organization_idB\x1f\n\x1d_counterparty_organization_idB\x12\n\x10_estimated_valueB\x15\n\x13_created_by_user_idB\r\n\x0b_created_atB\r\n\x0b_updated_at\"\xc3\x01\n\tTaskModel\x12\x0f\n\x07task_id\x18\x01 \x01(\x05\x12\x0f\n\x07lead_id\x18\x02 \x01(\x03\x12\x13\n\x0b\x61ssigned_to\x18\x03 \x01(\x03\x12\r\n\x05title\x18\x04 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x05 \x01(\t\x12\x0e\n\x06status\x18\x06 \x01(\t\x12\x11\n\tfrom_time\x18\x07 \x01(\t\x12\x10\n\x08\x64ue_time\x18\x08 \x01(\t\x12\x12\n\ncreated_at\x18\t \x01(\t\x12\x12\n\nupdated_at\x18\n \x01(\t\"z\n\tNoteModel\x12\x0f\n\x07note_id\x18\x01 \x01(\x03\x12\x0f\n\x07lead_id\x18\x02 \x01(\x03\x12\x12\n\ncreated_by\x18\x03 \x01(\x03\x12\x0f\n\x07\x63ontent\x18\x04 \x01(\t\x12\x12\n\ncreated_at\x18\x05 \x01(\t\x12\x12\n\nupdated_at\x18\x06 \x01(\t\"u\n\x10\x43hatMessageModel\x12\x12\n\nmessage_id\x18\x01 \x01(\x03\x12\x0f\n\x07lead_id\x18\x02 \x01(\x03\x12\x12\n\ncreated_by\x18\x03 \x01(\x03\x12\x14\n\x0cmessage_text\x18\x04 \x01(\t\x12\x12\n\ncreated_at\x18\x05 \x01(\t\"R\n\x0c\x41\x64\x64ressModel\x12\x0f\n\x02id\x18\x01 \x01(\x03H\x00\x88\x01\x01\x12\x19\n\x0c\x61\x64\x64ress_line\x18\x02 \x01(\tH\x01\x88\x01\x01\x42\x05\n\x03_idB\x0f\n\r_address_line\",\n\x11\x43reateLeadRequest\x12\x17\n\x04lead\x18\x01 \x01(\x0b\x32\t.crm.Lead\"-\n\x12\x43reateLeadResponse\x12\x17\n\x04lead\x18\x01 \x01(\x0b\x32\t.crm.Lead\"!\n\x0eGetLeadRequest\x12\x0f\n\x07lead_id\x18\x01 \x01(\x03\"*\n\x0fGetLeadResponse\x12\x17\n\x04lead\x18\x01 \x01(\x0b\x32\t.crm.Lead\"J\n\x10ListLeadsRequest\x12\x17\n\x0forganization_id\x18\x01 \x01(\x03\x12\r\n\x05limit\x18\x02 \x01(\x05\x12\x0e\n\x06offset\x18\x03 \x01(\x05\"B\n\x11ListLeadsResponse\x12\x18\n\x05leads\x18\x01 \x03(\x0b\x32\t.crm.Lead\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\",\n\x11UpdateLeadRequest\x12\x17\n\x04lead\x18\x01 \x01(\x0b\x32\t.crm.Lead\"-\n\x12UpdateLeadResponse\x12\x17\n\x04lead\x18\x01 \x01(\x0b\x32\t.crm.Lead\"+\n\x10PatchLeadRequest\x12\x17\n\x04lead\x18\x01 \x01(\x0b\x32\t.crm.Lead\",\n\x11PatchLeadResponse\x12\x17\n\x04lead\x18\x01 \x01(\x0b\x32\t.crm.Lead\"$\n\x11\x44\x65leteLeadRequest\x12\x0f\n\x07lead_id\x18\x01 \x01(\x03\"%\n\x12\x44\x65leteLeadResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\"\x92\x01\n\x11\x43reateTaskRequest\x12\x0f\n\x07lead_id\x18\x01 \x01(\x03\x12\x13\n\x0b\x61ssigned_to\x18\x02 \x01(\x03\x12\r\n\x05title\x18\x03 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x04 \x01(\t\x12\x0e\n\x06status\x18\x05 \x01(\t\x12\x11\n\tfrom_time\x18\x06 \x01(\t\x12\x10\n\x08\x64ue_time\x18\x07 \x01(\t\"2\n\x12\x43reateTaskResponse\x12\x1c\n\x04task\x18\x01 \x01(\x0b\x32\x0e.crm.TaskModel\"!\n\x0eGetTaskRequest\x12\x0f\n\x07task_id\x18\x01 \x01(\x03\"/\n\x0fGetTaskResponse\x12\x1c\n\x04task\x18\x01 \x01(\x0b\x32\x0e.crm.TaskModel\"D\n\x10ListTasksRequest\x12\x0f\n\x07lead_id\x18\x01 \x01(\x03\x12\x0c\n\x04page\x18\x02 \x01(\x05\x12\x11\n\tpage_size\x18\x03 \x01(\x05\"G\n\x11ListTasksResponse\x12\x1d\n\x05tasks\x18\x01 \x03(\x0b\x32\x0e.crm.TaskModel\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"\xa3\x01\n\x11UpdateTaskRequest\x12\x0f\n\x07task_id\x18\x01 \x01(\x03\x12\x0f\n\x07lead_id\x18\x02 \x01(\x03\x12\x13\n\x0b\x61ssigned_to\x18\x03 \x01(\x03\x12\r\n\x05title\x18\x04 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x05 \x01(\t\x12\x0e\n\x06status\x18\x06 \x01(\t\x12\x11\n\tfrom_time\x18\x07 \x01(\t\x12\x10\n\x08\x64ue_time\x18\x08 \x01(\t\"2\n\x12UpdateTaskResponse\x12\x1c\n\x04task\x18\x01 \x01(\x0b\x32\x0e.crm.TaskModel\"$\n\x11\x44\x65leteTaskRequest\x12\x0f\n\x07task_id\x18\x01 \x01(\x03\"%\n\x12\x44\x65leteTaskResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\"I\n\x11\x43reateNoteRequest\x12\x0f\n\x07lead_id\x18\x01 \x01(\x03\x12\x12\n\ncreated_by\x18\x02 \x01(\x03\x12\x0f\n\x07\x63ontent\x18\x03 \x01(\t\"2\n\x12\x43reateNoteResponse\x12\x1c\n\x04note\x18\x01 \x01(\x0b\x32\x0e.crm.NoteModel\"!\n\x0eGetNoteRequest\x12\x0f\n\x07note_id\x18\x01 \x01(\x03\"/\n\x0fGetNoteResponse\x12\x1c\n\x04note\x18\x01 \x01(\x0b\x32\x0e.crm.NoteModel\"D\n\x10ListNotesRequest\x12\x0f\n\x07lead_id\x18\x01 \x01(\x03\x12\x0c\n\x04page\x18\x02 \x01(\x05\x12\x11\n\tpage_size\x18\x03 \x01(\x05\"G\n\x11ListNotesResponse\x12\x1d\n\x05notes\x18\x01 \x03(\x0b\x32\x0e.crm.NoteModel\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"Z\n\x11UpdateNoteRequest\x12\x0f\n\x07note_id\x18\x01 \x01(\x03\x12\x0f\n\x07lead_id\x18\x02 \x01(\x03\x12\x12\n\ncreated_by\x18\x03 \x01(\x03\x12\x0f\n\x07\x63ontent\x18\x04 \x01(\t\"2\n\x12UpdateNoteResponse\x12\x1c\n\x04note\x18\x01 \x01(\x0b\x32\x0e.crm.NoteModel\"$\n\x11\x44\x65leteNoteRequest\x12\x0f\n\x07note_id\x18\x01 \x01(\x03\"%\n\x12\x44\x65leteNoteResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\"U\n\x18\x43reateChatMessageRequest\x12\x0f\n\x07lead_id\x18\x01 \x01(\x03\x12\x12\n\ncreated_by\x18\x02 \x01(\x03\x12\x14\n\x0cmessage_text\x18\x03 \x01(\t\"H\n\x19\x43reateChatMessageResponse\x12+\n\x0c\x63hat_message\x18\x01 \x01(\x0b\x32\x15.crm.ChatMessageModel\"+\n\x15GetChatMessageRequest\x12\x12\n\nmessage_id\x18\x01 \x01(\x03\"E\n\x16GetChatMessageResponse\x12+\n\x0c\x63hat_message\x18\x01 \x01(\x0b\x32\x15.crm.ChatMessageModel\"K\n\x17ListChatMessagesRequest\x12\x0f\n\x07lead_id\x18\x01 \x01(\x03\x12\x0c\n\x04page\x18\x02 \x01(\x05\x12\x11\n\tpage_size\x18\x03 \x01(\x05\"]\n\x18ListChatMessagesResponse\x12,\n\rchat_messages\x18\x01 \x03(\x0b\x32\x15.crm.ChatMessageModel\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"i\n\x18UpdateChatMessageRequest\x12\x12\n\nmessage_id\x18\x01 \x01(\x03\x12\x0f\n\x07lead_id\x18\x02 \x01(\x03\x12\x12\n\ncreated_by\x18\x03 \x01(\x03\x12\x14\n\x0cmessage_text\x18\x04 \x01(\t\"H\n\x19UpdateChatMessageResponse\x12+\n\x0c\x63hat_message\x18\x01 \x01(\x0b\x32\x15.crm.ChatMessageModel\".\n\x18\x44\x65leteChatMessageRequest\x12\x12\n\nmessage_id\x18\x01 \x01(\x03\",\n\x19\x44\x65leteChatMessageResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x32\xed\n\n\nCRMService\x12=\n\nCreateLead\x12\x16.crm.CreateLeadRequest\x1a\x17.crm.CreateLeadResponse\x12\x34\n\x07GetLead\x12\x13.crm.GetLeadRequest\x1a\x14.crm.GetLeadResponse\x12:\n\tListLeads\x12\x15.crm.ListLeadsRequest\x1a\x16.crm.ListLeadsResponse\x12=\n\nUpdateLead\x12\x16.crm.UpdateLeadRequest\x1a\x17.crm.UpdateLeadResponse\x12:\n\tPatchLead\x12\x15.crm.PatchLeadRequest\x1a\x16.crm.PatchLeadResponse\x12=\n\nDeleteLead\x12\x16.crm.DeleteLeadRequest\x1a\x17.crm.DeleteLeadResponse\x12=\n\nCreateTask\x12\x16.crm.CreateTaskRequest\x1a\x17.crm.CreateTaskResponse\x12\x34\n\x07GetTask\x12\x13.crm.GetTaskRequest\x1a\x14.crm.GetTaskResponse\x12:\n\tListTasks\x12\x15.crm.ListTasksRequest\x1a\x16.crm.ListTasksResponse\x12=\n\nUpdateTask\x12\x16.crm.UpdateTaskRequest\x1a\x17.crm.UpdateTaskResponse\x12=\n\nDeleteTask\x12\x16.crm.DeleteTaskRequest\x1a\x17.crm.DeleteTaskResponse\x12=\n\nCreateNote\x12\x16.crm.CreateNoteRequest\x1a\x17.crm.CreateNoteResponse\x12\x34\n\x07GetNote\x12\x13.crm.GetNoteRequest\x1a\x14.crm.GetNoteResponse\x12:\n\tListNotes\x12\x15.crm.ListNotesRequest\x1a\x16.crm.ListNotesResponse\x12=\n\nUpdateNote\x12\x16.crm.UpdateNoteRequest\x1a\x17.crm.UpdateNoteResponse\x12=\n\nDeleteNote\x12\x16.crm.DeleteNoteRequest\x1a\x17.crm.DeleteNoteResponse\x12R\n\x11\x43reateChatMessage\x12\x1d.crm.CreateChatMessageRequest\x1a\x1e.crm.CreateChatMessageResponse\x12I\n\x0eGetChatMessage\x12\x1a.crm.GetChatMessageRequest\x1a\x1b.crm.GetChatMessageResponse\x12O\n\x10ListChatMessages\x12\x1c.crm.ListChatMessagesRequest\x1a\x1d.crm.ListChatMessagesResponse\x12R\n\x11UpdateChatMessage\x12\x1d.crm.UpdateChatMessageRequest\x1a\x1e.crm.UpdateChatMessageResponse\x12R\n\x11\x44\x65leteChatMessage\x12\x1d.crm.DeleteChatMessageRequest\x1a\x1e.crm.DeleteChatMessageResponseB<Z:github.com/2group/2sales.core-service/pkg/gen/go/crm;crmv1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'crm.crm_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z:github.com/2group/2sales.core-service/pkg/gen/go/crm;crmv1'
  _globals['_ORGANIZATION']._serialized_start=40
  _globals['_ORGANIZATION']._serialized_end=208
  _globals['_LEAD']._serialized_start=211
  _globals['_LEAD']._serialized_end=689
  _globals['_CONTACT']._serialized_start=692
  _globals['_CONTACT']._serialized_end=947
  _globals['_LEADMODEL']._serialized_start=950
  _globals['_LEADMODEL']._serialized_end=1382
  _globals['_TASKMODEL']._serialized_start=1385
  _globals['_TASKMODEL']._serialized_end=1580
  _globals['_NOTEMODEL']._serialized_start=1582
  _globals['_NOTEMODEL']._serialized_end=1704
  _globals['_CHATMESSAGEMODEL']._serialized_start=1706
  _globals['_CHATMESSAGEMODEL']._serialized_end=1823
  _globals['_ADDRESSMODEL']._serialized_start=1825
  _globals['_ADDRESSMODEL']._serialized_end=1907
  _globals['_CREATELEADREQUEST']._serialized_start=1909
  _globals['_CREATELEADREQUEST']._serialized_end=1953
  _globals['_CREATELEADRESPONSE']._serialized_start=1955
  _globals['_CREATELEADRESPONSE']._serialized_end=2000
  _globals['_GETLEADREQUEST']._serialized_start=2002
  _globals['_GETLEADREQUEST']._serialized_end=2035
  _globals['_GETLEADRESPONSE']._serialized_start=2037
  _globals['_GETLEADRESPONSE']._serialized_end=2079
  _globals['_LISTLEADSREQUEST']._serialized_start=2081
  _globals['_LISTLEADSREQUEST']._serialized_end=2155
  _globals['_LISTLEADSRESPONSE']._serialized_start=2157
  _globals['_LISTLEADSRESPONSE']._serialized_end=2223
  _globals['_UPDATELEADREQUEST']._serialized_start=2225
  _globals['_UPDATELEADREQUEST']._serialized_end=2269
  _globals['_UPDATELEADRESPONSE']._serialized_start=2271
  _globals['_UPDATELEADRESPONSE']._serialized_end=2316
  _globals['_PATCHLEADREQUEST']._serialized_start=2318
  _globals['_PATCHLEADREQUEST']._serialized_end=2361
  _globals['_PATCHLEADRESPONSE']._serialized_start=2363
  _globals['_PATCHLEADRESPONSE']._serialized_end=2407
  _globals['_DELETELEADREQUEST']._serialized_start=2409
  _globals['_DELETELEADREQUEST']._serialized_end=2445
  _globals['_DELETELEADRESPONSE']._serialized_start=2447
  _globals['_DELETELEADRESPONSE']._serialized_end=2484
  _globals['_CREATETASKREQUEST']._serialized_start=2487
  _globals['_CREATETASKREQUEST']._serialized_end=2633
  _globals['_CREATETASKRESPONSE']._serialized_start=2635
  _globals['_CREATETASKRESPONSE']._serialized_end=2685
  _globals['_GETTASKREQUEST']._serialized_start=2687
  _globals['_GETTASKREQUEST']._serialized_end=2720
  _globals['_GETTASKRESPONSE']._serialized_start=2722
  _globals['_GETTASKRESPONSE']._serialized_end=2769
  _globals['_LISTTASKSREQUEST']._serialized_start=2771
  _globals['_LISTTASKSREQUEST']._serialized_end=2839
  _globals['_LISTTASKSRESPONSE']._serialized_start=2841
  _globals['_LISTTASKSRESPONSE']._serialized_end=2912
  _globals['_UPDATETASKREQUEST']._serialized_start=2915
  _globals['_UPDATETASKREQUEST']._serialized_end=3078
  _globals['_UPDATETASKRESPONSE']._serialized_start=3080
  _globals['_UPDATETASKRESPONSE']._serialized_end=3130
  _globals['_DELETETASKREQUEST']._serialized_start=3132
  _globals['_DELETETASKREQUEST']._serialized_end=3168
  _globals['_DELETETASKRESPONSE']._serialized_start=3170
  _globals['_DELETETASKRESPONSE']._serialized_end=3207
  _globals['_CREATENOTEREQUEST']._serialized_start=3209
  _globals['_CREATENOTEREQUEST']._serialized_end=3282
  _globals['_CREATENOTERESPONSE']._serialized_start=3284
  _globals['_CREATENOTERESPONSE']._serialized_end=3334
  _globals['_GETNOTEREQUEST']._serialized_start=3336
  _globals['_GETNOTEREQUEST']._serialized_end=3369
  _globals['_GETNOTERESPONSE']._serialized_start=3371
  _globals['_GETNOTERESPONSE']._serialized_end=3418
  _globals['_LISTNOTESREQUEST']._serialized_start=3420
  _globals['_LISTNOTESREQUEST']._serialized_end=3488
  _globals['_LISTNOTESRESPONSE']._serialized_start=3490
  _globals['_LISTNOTESRESPONSE']._serialized_end=3561
  _globals['_UPDATENOTEREQUEST']._serialized_start=3563
  _globals['_UPDATENOTEREQUEST']._serialized_end=3653
  _globals['_UPDATENOTERESPONSE']._serialized_start=3655
  _globals['_UPDATENOTERESPONSE']._serialized_end=3705
  _globals['_DELETENOTEREQUEST']._serialized_start=3707
  _globals['_DELETENOTEREQUEST']._serialized_end=3743
  _globals['_DELETENOTERESPONSE']._serialized_start=3745
  _globals['_DELETENOTERESPONSE']._serialized_end=3782
  _globals['_CREATECHATMESSAGEREQUEST']._serialized_start=3784
  _globals['_CREATECHATMESSAGEREQUEST']._serialized_end=3869
  _globals['_CREATECHATMESSAGERESPONSE']._serialized_start=3871
  _globals['_CREATECHATMESSAGERESPONSE']._serialized_end=3943
  _globals['_GETCHATMESSAGEREQUEST']._serialized_start=3945
  _globals['_GETCHATMESSAGEREQUEST']._serialized_end=3988
  _globals['_GETCHATMESSAGERESPONSE']._serialized_start=3990
  _globals['_GETCHATMESSAGERESPONSE']._serialized_end=4059
  _globals['_LISTCHATMESSAGESREQUEST']._serialized_start=4061
  _globals['_LISTCHATMESSAGESREQUEST']._serialized_end=4136
  _globals['_LISTCHATMESSAGESRESPONSE']._serialized_start=4138
  _globals['_LISTCHATMESSAGESRESPONSE']._serialized_end=4231
  _globals['_UPDATECHATMESSAGEREQUEST']._serialized_start=4233
  _globals['_UPDATECHATMESSAGEREQUEST']._serialized_end=4338
  _globals['_UPDATECHATMESSAGERESPONSE']._serialized_start=4340
  _globals['_UPDATECHATMESSAGERESPONSE']._serialized_end=4412
  _globals['_DELETECHATMESSAGEREQUEST']._serialized_start=4414
  _globals['_DELETECHATMESSAGEREQUEST']._serialized_end=4460
  _globals['_DELETECHATMESSAGERESPONSE']._serialized_start=4462
  _globals['_DELETECHATMESSAGERESPONSE']._serialized_end=4506
  _globals['_CRMSERVICE']._serialized_start=4509
  _globals['_CRMSERVICE']._serialized_end=5898
# @@protoc_insertion_point(module_scope)
