package class

type MethodDescriptor struct {
    parameterTypes  []*FieldType
    returnType      *FieldType
}

func parseMethodDescriptor(descriptor string) (*MethodDescriptor) {
    parser := newMemberDescriptorParser(descriptor)
    return parser.parse()
}

func (self *MethodDescriptor) ParameterTypes() ([]*FieldType) {
    return self.parameterTypes
}
// parameterCount()
func (self *MethodDescriptor) argCount() (uint) {
    return uint(len(self.parameterTypes))
}

func (self *MethodDescriptor) addParameterType(t *FieldType) {
    pLen := len(self.parameterTypes)
    if pLen == cap(self.parameterTypes) {
        self.parameterTypes = make([]*FieldType, pLen, pLen + 4)
    }

    self.parameterTypes = self.parameterTypes[:pLen + 1]
    self.parameterTypes[pLen] = t
}
