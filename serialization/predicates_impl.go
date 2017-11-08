package serialization

const PREDICATE_FACTORY_ID = -32

type predicate struct {
	id int32
}

func newPredicate(id int32) *predicate {
	return &predicate{id}
}

func (sp *predicate) ReadData(input DataInput) error {
	return nil
}

func (sp *predicate) WriteData(output DataOutput) {
}

func (*predicate) FactoryId() int32 {
	return PREDICATE_FACTORY_ID
}

func (p *predicate) ClassId() int32 {
	return p.id
}

type SqlPredicate struct {
	*predicate
	sql string
}

func NewSqlPredicate(sql string) *SqlPredicate {
	return &SqlPredicate{newPredicate(SQL_PREDICATE), sql}
}

func (sp *SqlPredicate) ReadData(input DataInput) error {
	var err error
	sp.sql, err = input.ReadUTF()
	return err
}

func (sp *SqlPredicate) WriteData(output DataOutput) {
	output.WriteUTF(sp.sql)
}

type AndPredicate struct {
	*predicate
	predicates []IPredicate
}

func NewAndPredicate(predicates []IPredicate) *AndPredicate {
	return &AndPredicate{newPredicate(AND_PREDICATE), predicates}
}

func (ap *AndPredicate) ReadData(input DataInput) error {
	length, err := input.ReadInt32()
	if err != nil {
		return err
	}
	ap.predicates = make([]IPredicate, 0)
	for i := 0; i < int(length); i++ {
		pred, err := input.ReadObject()
		if err != nil {
			return err
		}
		ap.predicates[i] = pred.(IPredicate)
	}
	return nil
}

func (ap *AndPredicate) WriteData(output DataOutput) {
	output.WriteInt32(int32(len(ap.predicates)))
	for _, pred := range ap.predicates {
		output.WriteObject(pred)
	}
}

type BetweenPredicate struct {
	*predicate
	field string
	from  interface{}
	to    interface{}
}

func NewBetweenPredicate(field string, from interface{}, to interface{}) *BetweenPredicate {
	return &BetweenPredicate{newPredicate(BETWEEN_PREDICATE), field, from, to}
}

func (bp *BetweenPredicate) ReadData(input DataInput) error {
	var err error
	bp.field, err = input.ReadUTF()
	if err != nil {
		return err
	}
	bp.to, err = input.ReadObject()
	if err != nil {
		return err
	}
	bp.from, err = input.ReadObject()

	return err
}

func (bp *BetweenPredicate) WriteData(output DataOutput) {
	output.WriteUTF(bp.field)
	output.WriteObject(bp.to)
	output.WriteObject(bp.from)

}

type EqualPredicate struct {
	*predicate
	field string
	value interface{}
}

func NewEqualPredicate(field string, value interface{}) *EqualPredicate {
	return &EqualPredicate{newPredicate(EQUAL_PREDICATE), field, value}
}

func (ep *EqualPredicate) ReadData(input DataInput) error {
	var err error
	ep.field, err = input.ReadUTF()
	if err != nil {
		return err
	}
	ep.value, err = input.ReadObject()

	return err
}

func (ep *EqualPredicate) WriteData(output DataOutput) {
	output.WriteUTF(ep.field)
	output.WriteObject(ep.value)
}

type GreaterLessPredicate struct {
	*predicate
	field string
	value interface{}
	equal bool
	less  bool
}

func NewGreaterLessPredicate(field string, value interface{}, equal bool, less bool) *GreaterLessPredicate {
	return &GreaterLessPredicate{newPredicate(GREATERLESS_PREDICATE), field, value, equal, less}
}

func (glp *GreaterLessPredicate) ReadData(input DataInput) error {
	var err error
	glp.field, err = input.ReadUTF()
	if err != nil {
		return err
	}
	glp.value, err = input.ReadObject()
	if err != nil {
		return err
	}
	glp.equal, err = input.ReadBool()
	if err != nil {
		return err
	}
	glp.less, err = input.ReadBool()
	return err
}

func (glp *GreaterLessPredicate) WriteData(output DataOutput) {
	output.WriteUTF(glp.field)
	output.WriteObject(glp.value)
	output.WriteBool(glp.equal)
	output.WriteBool(glp.less)
}

type LikePredicate struct {
	*predicate
	field string
	expr  string
}

func NewLikePredicate(field string, expr string) *LikePredicate {
	return &LikePredicate{newPredicate(LIKE_PREDICATE), field, expr}
}

func (lp *LikePredicate) ReadData(input DataInput) error {
	var err error
	lp.field, err = input.ReadUTF()
	if err != nil {
		return err
	}
	lp.expr, err = input.ReadUTF()
	return err
}

func (lp *LikePredicate) WriteData(output DataOutput) {
	output.WriteUTF(lp.field)
	output.WriteUTF(lp.expr)
}

type ILikePredicate struct {
	*LikePredicate
}

func NewILikePredicate(field string, expr string) *ILikePredicate {
	return &ILikePredicate{&LikePredicate{newPredicate(ILIKE_PREDICATE), field, expr}}
}

type InPredicate struct {
	*predicate
	field  string
	values []interface{}
}

func NewInPredicate(field string, values []interface{}) *InPredicate {
	return &InPredicate{newPredicate(IN_PREDICATE), field, values}
}

func (ip *InPredicate) ReadData(input DataInput) error {
	var err error
	ip.field, err = input.ReadUTF()
	if err != nil {
		return err
	}
	length, err := input.ReadInt32()
	if err != nil {
		return err
	}
	ip.values = make([]interface{}, length)
	for i := int32(0); i < length; i++ {
		ip.values[i], err = input.ReadObject()
		if err != nil {
			return err
		}
	}
	return nil
}

func (ip *InPredicate) WriteData(output DataOutput) {
	output.WriteUTF(ip.field)
	output.WriteInt32(int32(len(ip.values)))
	for _, value := range ip.values {
		output.WriteObject(value)
	}
}

type InstanceOfPredicate struct {
	*predicate
	className string
}

func NewInstanceOfPredicate(className string) *InstanceOfPredicate {
	return &InstanceOfPredicate{newPredicate(INSTANCEOF_PREDICATE), className}
}

func (iop *InstanceOfPredicate) ReadData(input DataInput) error {
	var err error
	iop.className, err = input.ReadUTF()
	return err
}

func (iop *InstanceOfPredicate) WriteData(output DataOutput) {
	output.WriteUTF(iop.className)
}

type NotEqualPredicate struct {
	*EqualPredicate
}

func NewNotEqualPredicate(field string, value interface{}) *NotEqualPredicate {
	return &NotEqualPredicate{&EqualPredicate{newPredicate(NOTEQUAL_PREDICATE), field, value}}
}

type NotPredicate struct {
	*predicate
	pred IPredicate
}

func NewNotPredicate(pred IPredicate) *NotPredicate {
	return &NotPredicate{newPredicate(NOT_PREDICATE), pred}
}

func (np *NotPredicate) ReadData(input DataInput) error {
	i, err := input.ReadObject()
	np.pred = i.(IPredicate)
	return err
}

func (np *NotPredicate) WriteData(output DataOutput) {
	output.WriteObject(np.pred)
}

type OrPredicate struct {
	*predicate
	predicates []IPredicate
}

func NewOrPredicate(predicates []IPredicate) *OrPredicate {
	return &OrPredicate{newPredicate(OR_PREDICATE), predicates}
}

func (or *OrPredicate) ReadData(input DataInput) error {
	var err error
	length, err := input.ReadInt32()
	if err != nil {
		return err
	}
	or.predicates = make([]IPredicate, 0)
	for i := 0; i < int(length); i++ {
		pred, err := input.ReadObject()
		if err != nil {
			return err
		}
		or.predicates[i] = pred.(IPredicate)
	}
	return err
}

func (or *OrPredicate) WriteData(output DataOutput) {
	output.WriteInt32(int32(len(or.predicates)))
	for _, pred := range or.predicates {
		output.WriteObject(pred)
	}
}

type RegexPredicate struct {
	*predicate
	field string
	regex string
}

func NewRegexPredicate(field string, regex string) *RegexPredicate {
	return &RegexPredicate{newPredicate(REGEX_PREDICATE), field, regex}
}

func (rp *RegexPredicate) ReadData(input DataInput) error {
	var err error
	rp.field, err = input.ReadUTF()
	if err != nil {
		return err
	}
	rp.regex, err = input.ReadUTF()
	return err
}

func (rp *RegexPredicate) WriteData(output DataOutput) {
	output.WriteUTF(rp.field)
	output.WriteUTF(rp.regex)
}

type FalsePredicate struct {
	*predicate
}

func NewFalsePredicate() *FalsePredicate {
	return &FalsePredicate{newPredicate(FALSE_PREDICATE)}
}
func (fp *FalsePredicate) ReadData(input DataInput) error {
	//Empty method
	return nil
}

func (fp *FalsePredicate) WriteData(output DataOutput) {
	//Empty method
}

type TruePredicate struct {
	*predicate
}

func NewTruePredicate() *TruePredicate {
	return &TruePredicate{newPredicate(TRUE_PREDICATE)}
}
func (tp *TruePredicate) ReadData(input DataInput) error {
	//Empty method
	return nil
}

func (tp *TruePredicate) WriteData(output DataOutput) {
	//Empty method
}