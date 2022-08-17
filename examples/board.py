"""Example Honeycomb Board with queries/annotations."""
import pulumi

from pulumi_honeycomb import (
    Board,
    BoardQueryArgs,
    DerivedColumn,
    Query,
    QueryAnnotation,
    get_query_specification,
)


_DATASET = "my-dataset"


hours_column = DerivedColumn(
    "duration_hours",
    alias="duration_hours",
    expression="DIV($duration_ms, 3600000)",
    description="Duration converted from milliseconds to hours",
    dataset=_DATASET,
)


# https://docs.honeycomb.io/api/query-specification/#fields-on-a-query-specification
query_spec = get_query_specification(
    time_range=24 * 60 * 60,  # 1 day in seconds
    calculations=[
        {
            "op": "MAX",
            "column": "duration_hours",
        },
    ],
    filter_combination="AND",
    filters=[
        {
            "column": "name",
            "op": "=",
            "value": "my_main_span",
        },
    ],
    breakdowns=["pipe", "client"],
)


query = Query(
    "query",
    dataset=_DATASET,
    query_json=query_spec.json,
)


query_annotation = QueryAnnotation(
    "query_annotation",
    dataset=_DATASET,
    name="Duration in Hours (last 24h)",
    query_id=query.id,
)


main_board = Board(
    "main_board",
    name="Staging",
    description="Key summaries.",
    style="visual",
    queries=[
        BoardQueryArgs(
            dataset=_DATASET,
            query_id=query.id,
            query_style="graph",
        ),
    ],
)
