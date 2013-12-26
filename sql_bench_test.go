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
		QueryAllFourTagsSegmentedParallel(10)
	}
}

func BenchmarkQueryAllFourTagsSegmentedParallel100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsSegmentedParallel(100)
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
		QueryAllFourTagsSegmentedGroupedParallel(10)
	}
}

func BenchmarkQueryAllFourTagsSegmentedGroupedParallel100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsSegmentedGroupedParallel(100)
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
		QueryAllFourTagsShardedParallel(10)
	}
}

func BenchmarkQueryAllFourTagsShardedParallel100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsShardedParallel(100)
	}
}

func BenchmarkQueryAllFourTagsShardedDistinct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsShardedDistinct()
	}
}

func BenchmarkQueryAllFourTagsShardedDistinctParallel10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsShardedDistinctParallel(10)
	}
}

func BenchmarkQueryAllFourTagsShardedDistinctParallel100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsShardedDistinctParallel(100)
	}
}

func BenchmarkQueryAllFourTagsShardedGrouped(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsShardedGrouped()
	}
}

func BenchmarkQueryAllFourTagsShardedGroupedParallel10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsShardedGroupedParallel(10)
	}
}

func BenchmarkQueryAllFourTagsShardedGroupedParallel100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllFourTagsShardedGroupedParallel(100)
	}
}

//

func BenchmarkQueryAllTenTagsSegmented(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsSegmented()
	}
}

func BenchmarkQueryAllTenTagsSegmentedParallel10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsSegmentedParallel(10)
	}
}

func BenchmarkQueryAllTenTagsSegmentedParallel100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsSegmentedParallel(100)
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
		QueryAllTenTagsSegmentedGroupedParallel(10)
	}
}

func BenchmarkQueryAllTenTagsSegmentedGroupedParallel100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsSegmentedGroupedParallel(100)
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

//

func BenchmarkQueryAllTenTagsSharded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsSharded()
	}
}

func BenchmarkQueryAllTenTagsShardedParallel10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsShardedParallel(10)
	}
}

func BenchmarkQueryAllTenTagsShardedParallel100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsShardedParallel(100)
	}
}

func BenchmarkQueryAllTenTagsShardedDistinct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsShardedDistinct()
	}
}

func BenchmarkQueryAllTenTagsShardedGrouped(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsShardedGrouped()
	}
}

func BenchmarkQueryAllTenTagsShardedGroupedParallel10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsShardedGroupedParallel(10)
	}
}

func BenchmarkQueryAllTenTagsShardedGroupedParallel100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryAllTenTagsShardedGroupedParallel(100)
	}
}