# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: gang_rape.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='gang_rape.proto',
  package='protocol',
  serialized_pb=_b('\n\x0fgang_rape.proto\x12\x08protocol\"\x12\n\x03Msg\x12\x0b\n\x03\x63md\x18\x01 \x02(\t\"\x1e\n\x0eSystem_Msg_S2C\x12\x0c\n\x04\x63ode\x18\x01 \x02(\x05\"\x12\n\x10System_Heart_C2S\"\'\n\x10System_Heart_S2C\x12\x13\n\x0bserver_time\x18\x01 \x02(\x03\"0\n\x12System_Message_S2C\x12\x1a\n\x03msg\x18\x01 \x02(\x0b\x32\r.protocol.Msg\"\x1e\n\x0e\x43\x65ll_Login_C2S\x12\x0c\n\x04name\x18\x01 \x02(\t\"\x1f\n\x0f\x43\x65ll_Logout_C2S\x12\x0c\n\x04name\x18\x01 \x02(\t')
)
_sym_db.RegisterFileDescriptor(DESCRIPTOR)




_MSG = _descriptor.Descriptor(
  name='Msg',
  full_name='protocol.Msg',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cmd', full_name='protocol.Msg.cmd', index=0,
      number=1, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=29,
  serialized_end=47,
)


_SYSTEM_MSG_S2C = _descriptor.Descriptor(
  name='System_Msg_S2C',
  full_name='protocol.System_Msg_S2C',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='protocol.System_Msg_S2C.code', index=0,
      number=1, type=5, cpp_type=1, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=49,
  serialized_end=79,
)


_SYSTEM_HEART_C2S = _descriptor.Descriptor(
  name='System_Heart_C2S',
  full_name='protocol.System_Heart_C2S',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=81,
  serialized_end=99,
)


_SYSTEM_HEART_S2C = _descriptor.Descriptor(
  name='System_Heart_S2C',
  full_name='protocol.System_Heart_S2C',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='server_time', full_name='protocol.System_Heart_S2C.server_time', index=0,
      number=1, type=3, cpp_type=2, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=101,
  serialized_end=140,
)


_SYSTEM_MESSAGE_S2C = _descriptor.Descriptor(
  name='System_Message_S2C',
  full_name='protocol.System_Message_S2C',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='msg', full_name='protocol.System_Message_S2C.msg', index=0,
      number=1, type=11, cpp_type=10, label=2,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=142,
  serialized_end=190,
)


_CELL_LOGIN_C2S = _descriptor.Descriptor(
  name='Cell_Login_C2S',
  full_name='protocol.Cell_Login_C2S',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='protocol.Cell_Login_C2S.name', index=0,
      number=1, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=192,
  serialized_end=222,
)


_CELL_LOGOUT_C2S = _descriptor.Descriptor(
  name='Cell_Logout_C2S',
  full_name='protocol.Cell_Logout_C2S',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='protocol.Cell_Logout_C2S.name', index=0,
      number=1, type=9, cpp_type=9, label=2,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=224,
  serialized_end=255,
)

_SYSTEM_MESSAGE_S2C.fields_by_name['msg'].message_type = _MSG
DESCRIPTOR.message_types_by_name['Msg'] = _MSG
DESCRIPTOR.message_types_by_name['System_Msg_S2C'] = _SYSTEM_MSG_S2C
DESCRIPTOR.message_types_by_name['System_Heart_C2S'] = _SYSTEM_HEART_C2S
DESCRIPTOR.message_types_by_name['System_Heart_S2C'] = _SYSTEM_HEART_S2C
DESCRIPTOR.message_types_by_name['System_Message_S2C'] = _SYSTEM_MESSAGE_S2C
DESCRIPTOR.message_types_by_name['Cell_Login_C2S'] = _CELL_LOGIN_C2S
DESCRIPTOR.message_types_by_name['Cell_Logout_C2S'] = _CELL_LOGOUT_C2S

Msg = _reflection.GeneratedProtocolMessageType('Msg', (_message.Message,), dict(
  DESCRIPTOR = _MSG,
  __module__ = 'gang_rape_pb2'
  # @@protoc_insertion_point(class_scope:protocol.Msg)
  ))
_sym_db.RegisterMessage(Msg)

System_Msg_S2C = _reflection.GeneratedProtocolMessageType('System_Msg_S2C', (_message.Message,), dict(
  DESCRIPTOR = _SYSTEM_MSG_S2C,
  __module__ = 'gang_rape_pb2'
  # @@protoc_insertion_point(class_scope:protocol.System_Msg_S2C)
  ))
_sym_db.RegisterMessage(System_Msg_S2C)

System_Heart_C2S = _reflection.GeneratedProtocolMessageType('System_Heart_C2S', (_message.Message,), dict(
  DESCRIPTOR = _SYSTEM_HEART_C2S,
  __module__ = 'gang_rape_pb2'
  # @@protoc_insertion_point(class_scope:protocol.System_Heart_C2S)
  ))
_sym_db.RegisterMessage(System_Heart_C2S)

System_Heart_S2C = _reflection.GeneratedProtocolMessageType('System_Heart_S2C', (_message.Message,), dict(
  DESCRIPTOR = _SYSTEM_HEART_S2C,
  __module__ = 'gang_rape_pb2'
  # @@protoc_insertion_point(class_scope:protocol.System_Heart_S2C)
  ))
_sym_db.RegisterMessage(System_Heart_S2C)

System_Message_S2C = _reflection.GeneratedProtocolMessageType('System_Message_S2C', (_message.Message,), dict(
  DESCRIPTOR = _SYSTEM_MESSAGE_S2C,
  __module__ = 'gang_rape_pb2'
  # @@protoc_insertion_point(class_scope:protocol.System_Message_S2C)
  ))
_sym_db.RegisterMessage(System_Message_S2C)

Cell_Login_C2S = _reflection.GeneratedProtocolMessageType('Cell_Login_C2S', (_message.Message,), dict(
  DESCRIPTOR = _CELL_LOGIN_C2S,
  __module__ = 'gang_rape_pb2'
  # @@protoc_insertion_point(class_scope:protocol.Cell_Login_C2S)
  ))
_sym_db.RegisterMessage(Cell_Login_C2S)

Cell_Logout_C2S = _reflection.GeneratedProtocolMessageType('Cell_Logout_C2S', (_message.Message,), dict(
  DESCRIPTOR = _CELL_LOGOUT_C2S,
  __module__ = 'gang_rape_pb2'
  # @@protoc_insertion_point(class_scope:protocol.Cell_Logout_C2S)
  ))
_sym_db.RegisterMessage(Cell_Logout_C2S)


# @@protoc_insertion_point(module_scope)