function windowOnLoad(data) {
    const DATE_COL = 2; // 2nd column, 0-indexed

    function getYear(text) {
        const match = /\d{2}\/\d{2}\/(\d{2})/.exec(text);
        return match ? '20' + match[1] : null; // TODO: fix for 2100 season
    }

    const yearSelect = document.getElementById('yearFilter');
    const populateYears = (table) => {
        const previous = yearSelect.value;
        const years = [
            ...new Set(
                table.column(DATE_COL).data().toArray().map(getYear).filter(Boolean),
            ),
        ]
            .sort()
            .reverse();
        console.log(table.column(DATE_COL).data().toArray(), years);

        yearSelect.length = 1; // keep only all years option
        years.forEach((y) => yearSelect.add(new Option(y, y)));
        yearSelect.value = years.includes(previous) ? previous : '';
    };

    const table = new DataTable('#recaps-table', {
        data: data,
        fixedColumns: {
            start: 1,
        },
        scrollY: 'calc(100vh - 11rem)',
        scrollX: true,
        pageLength: 50,
        columnControl: [['orderAsc', 'orderDesc', 'orderRemove', 'searchDropdown']],
        columns: [
            {
                title: 'Group Name',
                searchable: true,
            },
            { title: 'Division', searchable: true },
            { title: 'Date', searchable: true },
            { title: 'Comp', searchable: true },
            { title: 'MP 1 STT' },
            { title: 'MP 1 MUS' },
            { title: 'MP 1 Tot', className: 'light-gray' },
            { title: 'MP 2 STT' },
            { title: 'MP 2 MUS' },
            { title: 'MP 2 Tot', className: 'light-gray' },
            { title: 'MP Tot', className: 'gray' },
            { title: 'ME COMP' },
            { title: 'ME COMM' },
            { title: 'ME Tot', className: 'light-gray' },
            { title: 'VE COMP' },
            { title: 'VE COMM' },
            { title: 'VE Tot', className: 'light-gray' },
            { title: 'GE Tot', className: 'gray' },
            { title: 'VP COMP' },
            { title: 'VP EXC' },
            { title: 'VP Tot', className: 'light-gray' },
            { title: 'Sub Total' },
            { title: 'PEN' },
            { title: 'PEN Tot' },
            { title: 'Overall Band', className: 'dark-gray' },
            { title: 'AUX COMP' },
            { title: 'AUX ACH' },
            { title: 'AUX Tot', className: 'light-gray' },
            { title: 'AUX Adj Tot', className: 'dark-gray' },
            { title: 'PERC COMP' },
            { title: 'PERC ACH' },
            { title: 'PERC Tot', className: 'light-gray' },
            { title: 'PERC ADJ Tot', className: 'dark-gray' },
        ],
        initComplete: function() {
            populateYears(this.api());
        },
    });

    // Filter table based on year selection
    DataTable.ext.search.push(function(_settings, rowData) {
        // only one table, don't have to check which one

        const selected = yearSelect.value;
        return !selected || getYear(rowData[DATE_COL]) === selected;
    });

    yearSelect.addEventListener('change', () => table.draw());
}
