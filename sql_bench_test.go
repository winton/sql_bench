package main

import(
	// "fmt"
	"testing"
)

// func BenchmarkQueryAllSingleTag(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllSingleTag()
// 	}
// }

// func BenchmarkQueryAllSingleTagDistinct(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllSingleTagDistinct()
// 	}
// }

// func BenchmarkQueryAllSingleTagGrouped(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllSingleTagGrouped()
// 	}
// }

// func BenchmarkQueryAllSingleTagSegmented(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllSingleTagSegmented()
// 	}
// }

// func BenchmarkQueryAllSingleTagSegmentedDistinct(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllSingleTagSegmentedDistinct()
// 	}
// }

// func BenchmarkQueryAllSingleTagSegmentedGrouped(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllSingleTagSegmentedGrouped()
// 	}
// }

// func BenchmarkQueryAllTwoTags(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllTwoTags()
// 	}
// }

// func BenchmarkQueryAllTwoTagsDistinct(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllTwoTagsDistinct()
// 	}
// }

// func BenchmarkQueryAllTwoTagsGrouped(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllTwoTagsGrouped()
// 	}
// }

// func BenchmarkQueryAllTwoTagsSegmented(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllTwoTagsSegmented()
// 	}
// }

// func BenchmarkQueryAllTwoTagsSegmentedDistinct(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllTwoTagsSegmentedDistinct()
// 	}
// }

// func BenchmarkQueryAllTwoTagsSegmentedGrouped(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllTwoTagsSegmentedGrouped()
// 	}
// }


// func BenchmarkQueryAllFourTags(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllFourTags()
// 	}
// }

// func BenchmarkQueryAllFourTagsDistinct(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllFourTagsDistinct()
// 	}
// }

// func BenchmarkQueryAllFourTagsGrouped(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		QueryAllFourTagsGrouped()
// 	}
// }

func BenchmarkQueryAllFourTagsSegmented(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsSegmented()
	}
}

func BenchmarkQueryAllFourTagsSegmentedParallel10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsSegmentedParallel10()
	}
}

func BenchmarkQueryAllFourTagsSegmentedDistinct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsSegmentedDistinct()
	}
}

func BenchmarkQueryAllFourTagsSegmentedGrouped(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsSegmentedGrouped()
	}
}

func BenchmarkQueryAllFourTagsSegmentedGroupedParallel10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsSegmentedGroupedParallel10()
	}
}

func BenchmarkQueryAllFourTagsTwoSegments(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsTwoSegments()
	}
}

func BenchmarkQueryAllFourTagsTwoSegmentsGrouped(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsTwoSegmentsGrouped()
	}
}

//

func BenchmarkQueryAllFourTagsSharded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsSharded()
	}
}

func BenchmarkQueryAllFourTagsShardedParallel10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsShardedParallel10()
	}
}

func BenchmarkQueryAllFourTagsShardedDistinct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsShardedDistinct()
	}
}

func BenchmarkQueryAllFourTagsShardedDistinctParallel10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsShardedDistinctParallel10()
	}
}

func BenchmarkQueryAllFourTagsShardedGrouped(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsShardedGrouped()
	}
}

func BenchmarkQueryAllFourTagsShardedGroupedParallel10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsShardedGroupedParallel10()
	}
}

//

func BenchmarkQueryAllTenTagsSegmented(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsSegmented()
	}
}

func BenchmarkQueryAllTenTagsSegmentedDistinct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsSegmentedDistinct()
	}
}

func BenchmarkQueryAllTenTagsSegmentedGrouped(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsSegmentedGrouped()
	}
}

func BenchmarkQueryAllTenTagsSegmentedGroupedParallel10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsSegmentedGroupedParallel10()
	}
}

func BenchmarkQueryAllTenTagsTwoSegments(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsTwoSegments()
	}
}

func BenchmarkQueryAllTenTagsTwoSegmentsGrouped(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsTwoSegmentsGrouped()
	}
}