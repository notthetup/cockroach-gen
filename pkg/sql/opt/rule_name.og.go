// Code generated by optgen; DO NOT EDIT.

package opt

const (
	startAutoRule RuleName = iota + NumManualRuleNames

	// ------------------------------------------------------------
	// Normalize Rule Names
	// ------------------------------------------------------------
	EliminateEmptyAnd
	EliminateEmptyOr
	EliminateSingletonAndOr
	SimplifyAnd
	SimplifyOr
	SimplifyFilters
	FoldNullAndOr
	NegateComparison
	EliminateNot
	NegateAnd
	NegateOr
	ExtractRedundantClause
	ExtractRedundantSubclause
	CommuteVarInequality
	CommuteConstInequality
	NormalizeCmpPlusConst
	NormalizeCmpMinusConst
	NormalizeCmpConstMinus
	NormalizeTupleEquality
	FoldNullComparisonLeft
	FoldNullComparisonRight
	FoldIsNull
	FoldNonNullIsNull
	FoldIsNotNull
	FoldNonNullIsNotNull
	CommuteNullIs
	DecorrelateJoin
	TryDecorrelateSelect
	TryDecorrelateProject
	TryDecorrelateProjectSelect
	TryDecorrelateScalarGroupBy
	HoistSelectExists
	HoistSelectNotExists
	HoistSelectSubquery
	HoistProjectSubquery
	HoistJoinSubquery
	HoistValuesSubquery
	NormalizeAnyFilter
	NormalizeNotAnyFilter
	EliminateDistinct
	EliminateGroupByProject
	PushSelectIntoInlinableProject
	InlineProjectInProject
	EnsureJoinFiltersAnd
	EnsureJoinFilters
	PushFilterIntoJoinLeft
	PushFilterIntoJoinRight
	SimplifyLeftJoin
	SimplifyRightJoin
	EliminateSemiJoin
	EliminateAntiJoin
	EliminateJoinNoColsLeft
	EliminateJoinNoColsRight
	EliminateLimit
	PushLimitIntoProject
	PushOffsetIntoProject
	EliminateMax1Row
	FoldPlusZero
	FoldZeroPlus
	FoldMinusZero
	FoldMultOne
	FoldOneMult
	FoldDivOne
	InvertMinus
	EliminateUnaryMinus
	EliminateProject
	EliminateProjectProject
	PruneProjectCols
	PruneScanCols
	PruneSelectCols
	PruneLimitCols
	PruneOffsetCols
	PruneJoinLeftCols
	PruneJoinRightCols
	PruneAggCols
	PruneGroupByCols
	PruneValuesCols
	PruneRowNumberCols
	CommuteVar
	CommuteConst
	EliminateCoalesce
	SimplifyCoalesce
	EliminateCast
	FoldNullCast
	FoldNullUnary
	FoldNullBinaryLeft
	FoldNullBinaryRight
	FoldNullInNonEmpty
	FoldNullInEmpty
	FoldNullNotInEmpty
	NormalizeInConst
	FoldInNull
	EliminateExistsProject
	EliminateExistsGroupBy
	EliminateSelect
	EnsureSelectFiltersAnd
	EnsureSelectFilters
	MergeSelects
	PushSelectIntoProject
	PushSelectIntoJoinLeft
	PushSelectIntoJoinRight
	MergeSelectInnerJoin
	PushSelectIntoGroupBy

	// ------------------------------------------------------------
	// Explore Rule Names
	// ------------------------------------------------------------
	PushLimitIntoScan
	PushLimitIntoLookupJoin
	GenerateIndexScans
	ConstrainScan
	PushFilterIntoLookupJoinNoRemainder
	PushFilterIntoLookupJoin
	ConstrainLookupJoinIndexScan

	// NumRuleNames tracks the total count of rule names.
	NumRuleNames
)
