// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: roachpb/app_stats.proto

package roachpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type StatementStatistics struct {
	// Count is the total number of times this statement was executed
	// since the begin of the reporting period.
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count"`
	// FirstAttemptCount collects the total number of times a first
	// attempt was executed (either the one time in explicitly committed
	// statements, or the first time in implicitly committed statements
	// with implicit retries).
	// The proportion of statements that could be executed without retry
	// can be computed as FirstAttemptCount / Count.
	// The cumulative number of retries can be computed with
	// Count - FirstAttemptCount.
	FirstAttemptCount int64 `protobuf:"varint,2,opt,name=first_attempt_count,json=firstAttemptCount" json:"first_attempt_count"`
	// MaxRetries collects the maximum observed number of automatic
	// retries in the reporting period.
	MaxRetries int64 `protobuf:"varint,3,opt,name=max_retries,json=maxRetries" json:"max_retries"`
	// LastErr collects the last error encountered.
	LastErr string `protobuf:"bytes,4,opt,name=last_err,json=lastErr" json:"last_err"`
	// LastErrRedacted collects the last error, redacted for reporting.
	LastErrRedacted string `protobuf:"bytes,11,opt,name=last_err_redacted,json=lastErrRedacted" json:"last_err_redacted"`
	// NumRows collects the number of rows returned or observed.
	NumRows NumericStat `protobuf:"bytes,5,opt,name=num_rows,json=numRows" json:"num_rows"`
	// ParseLat is the time to transform the SQL string into an AST.
	ParseLat NumericStat `protobuf:"bytes,6,opt,name=parse_lat,json=parseLat" json:"parse_lat"`
	// PlanLat is the time to transform the AST into a logical query plan.
	PlanLat NumericStat `protobuf:"bytes,7,opt,name=plan_lat,json=planLat" json:"plan_lat"`
	// RunLat is the time to run the query and fetch/compute the result rows.
	RunLat NumericStat `protobuf:"bytes,8,opt,name=run_lat,json=runLat" json:"run_lat"`
	// ServiceLat is the time to service the query, from start of parse to end of execute.
	ServiceLat NumericStat `protobuf:"bytes,9,opt,name=service_lat,json=serviceLat" json:"service_lat"`
	// OverheadLat is the difference between ServiceLat and the sum of parse+plan+run latencies.
	// We store it separately (as opposed to computing it post-hoc) because the combined
	// variance for the overhead cannot be derived from the variance of the separate latencies.
	OverheadLat NumericStat `protobuf:"bytes,10,opt,name=overhead_lat,json=overheadLat" json:"overhead_lat"`
}

func (m *StatementStatistics) Reset()                    { *m = StatementStatistics{} }
func (m *StatementStatistics) String() string            { return proto.CompactTextString(m) }
func (*StatementStatistics) ProtoMessage()               {}
func (*StatementStatistics) Descriptor() ([]byte, []int) { return fileDescriptorAppStats, []int{0} }

type NumericStat struct {
	// NumericStat keeps track of two running values --- the running mean and
	// the running sum of squared differences from the mean. Using this along
	// with the total count of values, we can compute variance using Welford's
	// method. This is more reliable than keeping track of the sum of
	// squared values, which is liable to overflow. See
	// https://en.wikipedia.org/wiki/Algorithms_for_calculating_variance#Online_algorithm
	Mean         float64 `protobuf:"fixed64,1,opt,name=mean" json:"mean"`
	SquaredDiffs float64 `protobuf:"fixed64,2,opt,name=squared_diffs,json=squaredDiffs" json:"squared_diffs"`
}

func (m *NumericStat) Reset()                    { *m = NumericStat{} }
func (m *NumericStat) String() string            { return proto.CompactTextString(m) }
func (*NumericStat) ProtoMessage()               {}
func (*NumericStat) Descriptor() ([]byte, []int) { return fileDescriptorAppStats, []int{1} }

type StatementStatisticsKey struct {
	Query   string `protobuf:"bytes,1,opt,name=query" json:"query"`
	App     string `protobuf:"bytes,2,opt,name=app" json:"app"`
	DistSQL bool   `protobuf:"varint,3,opt,name=distSQL" json:"distSQL"`
	Failed  bool   `protobuf:"varint,4,opt,name=failed" json:"failed"`
	Opt     bool   `protobuf:"varint,5,opt,name=opt" json:"opt"`
}

func (m *StatementStatisticsKey) Reset()                    { *m = StatementStatisticsKey{} }
func (m *StatementStatisticsKey) String() string            { return proto.CompactTextString(m) }
func (*StatementStatisticsKey) ProtoMessage()               {}
func (*StatementStatisticsKey) Descriptor() ([]byte, []int) { return fileDescriptorAppStats, []int{2} }

// CollectedStats wraps collected timings and metadata for some query's execution.
type CollectedStatementStatistics struct {
	Key   StatementStatisticsKey `protobuf:"bytes,1,opt,name=key" json:"key"`
	Stats StatementStatistics    `protobuf:"bytes,2,opt,name=stats" json:"stats"`
}

func (m *CollectedStatementStatistics) Reset()         { *m = CollectedStatementStatistics{} }
func (m *CollectedStatementStatistics) String() string { return proto.CompactTextString(m) }
func (*CollectedStatementStatistics) ProtoMessage()    {}
func (*CollectedStatementStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptorAppStats, []int{3}
}

func init() {
	proto.RegisterType((*StatementStatistics)(nil), "cockroach.sql.StatementStatistics")
	proto.RegisterType((*NumericStat)(nil), "cockroach.sql.NumericStat")
	proto.RegisterType((*StatementStatisticsKey)(nil), "cockroach.sql.StatementStatisticsKey")
	proto.RegisterType((*CollectedStatementStatistics)(nil), "cockroach.sql.CollectedStatementStatistics")
}
func (m *StatementStatistics) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatementStatistics) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintAppStats(dAtA, i, uint64(m.Count))
	dAtA[i] = 0x10
	i++
	i = encodeVarintAppStats(dAtA, i, uint64(m.FirstAttemptCount))
	dAtA[i] = 0x18
	i++
	i = encodeVarintAppStats(dAtA, i, uint64(m.MaxRetries))
	dAtA[i] = 0x22
	i++
	i = encodeVarintAppStats(dAtA, i, uint64(len(m.LastErr)))
	i += copy(dAtA[i:], m.LastErr)
	dAtA[i] = 0x2a
	i++
	i = encodeVarintAppStats(dAtA, i, uint64(m.NumRows.Size()))
	n1, err := m.NumRows.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x32
	i++
	i = encodeVarintAppStats(dAtA, i, uint64(m.ParseLat.Size()))
	n2, err := m.ParseLat.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x3a
	i++
	i = encodeVarintAppStats(dAtA, i, uint64(m.PlanLat.Size()))
	n3, err := m.PlanLat.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x42
	i++
	i = encodeVarintAppStats(dAtA, i, uint64(m.RunLat.Size()))
	n4, err := m.RunLat.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x4a
	i++
	i = encodeVarintAppStats(dAtA, i, uint64(m.ServiceLat.Size()))
	n5, err := m.ServiceLat.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	dAtA[i] = 0x52
	i++
	i = encodeVarintAppStats(dAtA, i, uint64(m.OverheadLat.Size()))
	n6, err := m.OverheadLat.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	dAtA[i] = 0x5a
	i++
	i = encodeVarintAppStats(dAtA, i, uint64(len(m.LastErrRedacted)))
	i += copy(dAtA[i:], m.LastErrRedacted)
	return i, nil
}

func (m *NumericStat) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NumericStat) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x9
	i++
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Mean))))
	i += 8
	dAtA[i] = 0x11
	i++
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.SquaredDiffs))))
	i += 8
	return i, nil
}

func (m *StatementStatisticsKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatementStatisticsKey) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintAppStats(dAtA, i, uint64(len(m.Query)))
	i += copy(dAtA[i:], m.Query)
	dAtA[i] = 0x12
	i++
	i = encodeVarintAppStats(dAtA, i, uint64(len(m.App)))
	i += copy(dAtA[i:], m.App)
	dAtA[i] = 0x18
	i++
	if m.DistSQL {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x20
	i++
	if m.Failed {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	dAtA[i] = 0x28
	i++
	if m.Opt {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}

func (m *CollectedStatementStatistics) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CollectedStatementStatistics) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintAppStats(dAtA, i, uint64(m.Key.Size()))
	n7, err := m.Key.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	dAtA[i] = 0x12
	i++
	i = encodeVarintAppStats(dAtA, i, uint64(m.Stats.Size()))
	n8, err := m.Stats.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	return i, nil
}

func encodeVarintAppStats(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *StatementStatistics) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovAppStats(uint64(m.Count))
	n += 1 + sovAppStats(uint64(m.FirstAttemptCount))
	n += 1 + sovAppStats(uint64(m.MaxRetries))
	l = len(m.LastErr)
	n += 1 + l + sovAppStats(uint64(l))
	l = m.NumRows.Size()
	n += 1 + l + sovAppStats(uint64(l))
	l = m.ParseLat.Size()
	n += 1 + l + sovAppStats(uint64(l))
	l = m.PlanLat.Size()
	n += 1 + l + sovAppStats(uint64(l))
	l = m.RunLat.Size()
	n += 1 + l + sovAppStats(uint64(l))
	l = m.ServiceLat.Size()
	n += 1 + l + sovAppStats(uint64(l))
	l = m.OverheadLat.Size()
	n += 1 + l + sovAppStats(uint64(l))
	l = len(m.LastErrRedacted)
	n += 1 + l + sovAppStats(uint64(l))
	return n
}

func (m *NumericStat) Size() (n int) {
	var l int
	_ = l
	n += 9
	n += 9
	return n
}

func (m *StatementStatisticsKey) Size() (n int) {
	var l int
	_ = l
	l = len(m.Query)
	n += 1 + l + sovAppStats(uint64(l))
	l = len(m.App)
	n += 1 + l + sovAppStats(uint64(l))
	n += 2
	n += 2
	n += 2
	return n
}

func (m *CollectedStatementStatistics) Size() (n int) {
	var l int
	_ = l
	l = m.Key.Size()
	n += 1 + l + sovAppStats(uint64(l))
	l = m.Stats.Size()
	n += 1 + l + sovAppStats(uint64(l))
	return n
}

func sovAppStats(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAppStats(x uint64) (n int) {
	return sovAppStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StatementStatistics) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAppStats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatementStatistics: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatementStatistics: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FirstAttemptCount", wireType)
			}
			m.FirstAttemptCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FirstAttemptCount |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRetries", wireType)
			}
			m.MaxRetries = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRetries |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastErr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAppStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastErr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumRows", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAppStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NumRows.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParseLat", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAppStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ParseLat.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanLat", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAppStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PlanLat.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunLat", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAppStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RunLat.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceLat", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAppStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ServiceLat.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OverheadLat", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAppStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OverheadLat.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastErrRedacted", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAppStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastErrRedacted = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAppStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAppStats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NumericStat) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAppStats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NumericStat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NumericStat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mean", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Mean = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field SquaredDiffs", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.SquaredDiffs = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipAppStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAppStats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatementStatisticsKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAppStats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StatementStatisticsKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatementStatisticsKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAppStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field App", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAppStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.App = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistSQL", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DistSQL = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Failed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Failed = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Opt", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Opt = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAppStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAppStats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CollectedStatementStatistics) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAppStats
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CollectedStatementStatistics: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CollectedStatementStatistics: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAppStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Key.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAppStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAppStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAppStats
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAppStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAppStats
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAppStats
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthAppStats
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAppStats
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAppStats(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAppStats = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAppStats   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("roachpb/app_stats.proto", fileDescriptorAppStats) }

var fileDescriptorAppStats = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xbb, 0xe4, 0xff, 0xb8, 0x15, 0xea, 0x16, 0x15, 0x2b, 0xaa, 0xdc, 0x10, 0xa9, 0x52,
	0xb8, 0xa4, 0x28, 0xe2, 0x82, 0x50, 0x91, 0xda, 0xc0, 0x89, 0x0a, 0x89, 0xf4, 0xc6, 0xc5, 0x5a,
	0xec, 0x49, 0x6b, 0xc5, 0xf6, 0x6e, 0x76, 0xd7, 0x6d, 0xf3, 0x16, 0xbc, 0x00, 0x17, 0x9e, 0x26,
	0x47, 0x8e, 0x5c, 0x40, 0x10, 0x5e, 0x04, 0xed, 0xda, 0x2b, 0x05, 0x14, 0x89, 0xdc, 0xac, 0xef,
	0x9b, 0xdf, 0xb7, 0x3b, 0x9e, 0xb1, 0xe1, 0xb1, 0xe4, 0x2c, 0xba, 0x11, 0x1f, 0x4f, 0x99, 0x10,
	0xa1, 0xd2, 0x4c, 0xab, 0xa1, 0x90, 0x5c, 0x73, 0xba, 0x17, 0xf1, 0x68, 0x66, 0xcd, 0xa1, 0x9a,
	0xa7, 0xdd, 0x47, 0xd7, 0xfc, 0x9a, 0x5b, 0xe7, 0xd4, 0x3c, 0x95, 0x45, 0xfd, 0xef, 0x75, 0x38,
	0xb8, 0xd2, 0x4c, 0x63, 0x86, 0xb9, 0x36, 0x0f, 0x89, 0xd2, 0x49, 0xa4, 0x68, 0x17, 0x1a, 0x11,
	0x2f, 0x72, 0xed, 0x93, 0x1e, 0x19, 0xd4, 0x2e, 0xea, 0xcb, 0x1f, 0xc7, 0x3b, 0x93, 0x52, 0xa2,
	0xcf, 0xe1, 0x60, 0x9a, 0x48, 0xa5, 0x43, 0xa6, 0x35, 0x66, 0x42, 0x87, 0x65, 0xe5, 0x83, 0xb5,
	0xca, 0x7d, 0x5b, 0x70, 0x5e, 0xfa, 0x63, 0x4b, 0x9d, 0x80, 0x97, 0xb1, 0xfb, 0x50, 0xa2, 0x96,
	0x09, 0x2a, 0xbf, 0xb6, 0x56, 0x0d, 0x19, 0xbb, 0x9f, 0x94, 0x3a, 0x3d, 0x86, 0x76, 0xca, 0x94,
	0x0e, 0x51, 0x4a, 0xbf, 0xde, 0x23, 0x83, 0x4e, 0x55, 0xd3, 0x32, 0xea, 0x1b, 0x29, 0xe9, 0x4b,
	0x68, 0xe7, 0x45, 0x16, 0x4a, 0x7e, 0xa7, 0xfc, 0x46, 0x8f, 0x0c, 0xbc, 0x51, 0x77, 0xf8, 0x57,
	0xa7, 0xc3, 0x77, 0x45, 0x86, 0x32, 0x89, 0x4c, 0x37, 0x0e, 0xce, 0x8b, 0x6c, 0xc2, 0xef, 0x14,
	0x3d, 0x83, 0x8e, 0x60, 0x52, 0x61, 0x98, 0x32, 0xed, 0x37, 0xb7, 0xa4, 0xdb, 0x16, 0xb9, 0x64,
	0xda, 0x9c, 0x2d, 0x52, 0x96, 0x5b, 0xba, 0xb5, 0xed, 0xd9, 0x86, 0x30, 0xf0, 0x0b, 0x68, 0xc9,
	0xa2, 0x64, 0xdb, 0x5b, 0xb2, 0x4d, 0x59, 0x58, 0xf4, 0x1c, 0x3c, 0x85, 0xf2, 0x36, 0x89, 0xca,
	0x8b, 0x77, 0xb6, 0xc4, 0xa1, 0x82, 0x4c, 0xc4, 0x18, 0x76, 0xf9, 0x2d, 0xca, 0x1b, 0x64, 0xb1,
	0xcd, 0x80, 0x2d, 0x33, 0x3c, 0x47, 0x99, 0x90, 0x67, 0xb0, 0xef, 0x86, 0x13, 0x4a, 0x8c, 0x59,
	0xa4, 0x31, 0xf6, 0xbd, 0xb5, 0x29, 0x3d, 0xac, 0xa6, 0x34, 0xa9, 0xcc, 0xfe, 0x04, 0xbc, 0xb5,
	0x4c, 0xea, 0x43, 0x3d, 0x43, 0x96, 0xdb, 0xad, 0x22, 0x15, 0x63, 0x15, 0xfa, 0x14, 0xf6, 0xd4,
	0xbc, 0x60, 0x12, 0xe3, 0x30, 0x4e, 0xa6, 0x53, 0x65, 0xd7, 0xc9, 0x95, 0xec, 0x56, 0xd6, 0x6b,
	0xe3, 0xf4, 0xbf, 0x10, 0x38, 0xdc, 0xb0, 0xb3, 0x6f, 0x71, 0x61, 0xd6, 0x76, 0x5e, 0xa0, 0x5c,
	0xd8, 0x03, 0xdc, 0xa5, 0x4a, 0x89, 0x1e, 0x42, 0x8d, 0x09, 0x61, 0x73, 0x9d, 0x63, 0x04, 0x1a,
	0x40, 0x2b, 0x4e, 0x94, 0xbe, 0x7a, 0x7f, 0x69, 0x97, 0xb2, 0xed, 0xe6, 0x56, 0x89, 0xf4, 0x08,
	0x9a, 0x53, 0x96, 0xa4, 0x18, 0xdb, 0x7d, 0x74, 0x76, 0xa5, 0x99, 0x54, 0x2e, 0xb4, 0xdd, 0x44,
	0x67, 0x19, 0xa1, 0xff, 0x99, 0xc0, 0xd1, 0x98, 0xa7, 0x29, 0x9a, 0xd7, 0xb0, 0xe9, 0x0b, 0x3b,
	0x83, 0xda, 0x0c, 0xcb, 0x8b, 0x7a, 0xa3, 0x93, 0x7f, 0xe6, 0xb0, 0xb9, 0x3d, 0x97, 0x3f, 0xc3,
	0x05, 0x7d, 0x05, 0x0d, 0xfb, 0xb1, 0xdb, 0x7e, 0xbc, 0x51, 0xff, 0xff, 0x01, 0xee, 0x6d, 0x58,
	0xec, 0xe2, 0xc9, 0xf2, 0x57, 0xb0, 0xb3, 0x5c, 0x05, 0xe4, 0xeb, 0x2a, 0x20, 0xdf, 0x56, 0x01,
	0xf9, 0xb9, 0x0a, 0xc8, 0xa7, 0xdf, 0xc1, 0xce, 0x87, 0x56, 0xf5, 0x43, 0xf9, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x4d, 0xe7, 0x67, 0x54, 0x5a, 0x04, 0x00, 0x00,
}
