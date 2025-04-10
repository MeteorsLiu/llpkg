package liblzma

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const LZMA_VERSION_MAJOR = 5
const LZMA_VERSION_MINOR = 4
const LZMA_VERSION_PATCH = 5
const LZMA_VERSION_COMMIT = ""
const LZMA_VERSION_STABILITY_ALPHA = 0
const LZMA_VERSION_STABILITY_BETA = 1
const LZMA_VERSION_STABILITY_STABLE = 2
const LZMA_VERSION_STABILITY_STRING = ""
const LZMA_VLI_BYTES_MAX = 9
const LZMA_CHECK_ID_MAX = 15
const LZMA_CHECK_SIZE_MAX = 64
const LZMA_FILTERS_MAX = 4
const LZMA_DELTA_DIST_MIN = 1
const LZMA_DELTA_DIST_MAX = 256
const LZMA_LCLP_MIN = 0
const LZMA_LCLP_MAX = 4
const LZMA_LC_DEFAULT = 3
const LZMA_LP_DEFAULT = 0
const LZMA_PB_MIN = 0
const LZMA_PB_MAX = 4
const LZMA_PB_DEFAULT = 2
const LZMA_STREAM_HEADER_SIZE = 12
const LZMA_BACKWARD_SIZE_MIN = 4
const LZMA_BLOCK_HEADER_SIZE_MIN = 8
const LZMA_BLOCK_HEADER_SIZE_MAX = 1024

/**
 * \brief       Run-time version number as an integer
 *
 * This allows an application to compare if it was built against the same,
 * older, or newer version of liblzma that is currently running.
 *
 * \return The value of LZMA_VERSION macro at the compile time of liblzma
 */
//go:linkname LzmaVersionNumber C.lzma_version_number
func LzmaVersionNumber() uint32

/**
 * \brief       Run-time version as a string
 *
 * This function may be useful to display which version of liblzma an
 * application is currently using.
 *
 * \return      Run-time version of liblzma
 */
//go:linkname LzmaVersionString C.lzma_version_string
func LzmaVersionString() *int8

type LzmaBool int8
type LzmaReservedEnum c.Int

const LZMARESERVEDENUM LzmaReservedEnum = 0

type LzmaRet c.Int

const (
	LZMAOK               LzmaRet = 0
	LZMASTREAMEND        LzmaRet = 1
	LZMANOCHECK          LzmaRet = 2
	LZMAUNSUPPORTEDCHECK LzmaRet = 3
	LZMAGETCHECK         LzmaRet = 4
	LZMAMEMERROR         LzmaRet = 5
	LZMAMEMLIMITERROR    LzmaRet = 6
	LZMAFORMATERROR      LzmaRet = 7
	LZMAOPTIONSERROR     LzmaRet = 8
	LZMADATAERROR        LzmaRet = 9
	LZMABUFERROR         LzmaRet = 10
	LZMAPROGERROR        LzmaRet = 11
	LZMASEEKNEEDED       LzmaRet = 12
	LZMARETINTERNAL1     LzmaRet = 101
	LZMARETINTERNAL2     LzmaRet = 102
	LZMARETINTERNAL3     LzmaRet = 103
	LZMARETINTERNAL4     LzmaRet = 104
	LZMARETINTERNAL5     LzmaRet = 105
	LZMARETINTERNAL6     LzmaRet = 106
	LZMARETINTERNAL7     LzmaRet = 107
	LZMARETINTERNAL8     LzmaRet = 108
)

type LzmaAction c.Int

const (
	LZMARUN         LzmaAction = 0
	LZMASYNCFLUSH   LzmaAction = 1
	LZMAFULLFLUSH   LzmaAction = 2
	LZMAFULLBARRIER LzmaAction = 4
	LZMAFINISH      LzmaAction = 3
)

/**
 * \brief       Custom functions for memory handling
 *
 * A pointer to lzma_allocator may be passed via lzma_stream structure
 * to liblzma, and some advanced functions take a pointer to lzma_allocator
 * as a separate function argument. The library will use the functions
 * specified in lzma_allocator for memory handling instead of the default
 * malloc() and free(). C++ users should note that the custom memory
 * handling functions must not throw exceptions.
 *
 * Single-threaded mode only: liblzma doesn't make an internal copy of
 * lzma_allocator. Thus, it is OK to change these function pointers in
 * the middle of the coding process, but obviously it must be done
 * carefully to make sure that the replacement `free' can deallocate
 * memory allocated by the earlier `alloc' function(s).
 *
 * Multithreaded mode: liblzma might internally store pointers to the
 * lzma_allocator given via the lzma_stream structure. The application
 * must not change the allocator pointer in lzma_stream or the contents
 * of the pointed lzma_allocator structure until lzma_end() has been used
 * to free the memory associated with that lzma_stream. The allocation
 * functions might be called simultaneously from multiple threads, and
 * thus they must be thread safe.
 */

type LzmaAllocator struct {
	Alloc  unsafe.Pointer
	Free   unsafe.Pointer
	Opaque unsafe.Pointer
}

type LzmaInternalS struct {
	Unused [8]uint8
}
type LzmaInternal LzmaInternalS

/**
 * \brief       Passing data to and from liblzma
 *
 * The lzma_stream structure is used for
 *  - passing pointers to input and output buffers to liblzma;
 *  - defining custom memory handler functions; and
 *  - holding a pointer to coder-specific internal data structures.
 *
 * Typical usage:
 *
 *  - After allocating lzma_stream (on stack or with malloc()), it must be
 *    initialized to LZMA_STREAM_INIT (see LZMA_STREAM_INIT for details).
 *
 *  - Initialize a coder to the lzma_stream, for example by using
 *    lzma_easy_encoder() or lzma_auto_decoder(). Some notes:
 *      - In contrast to zlib, strm->next_in and strm->next_out are
 *        ignored by all initialization functions, thus it is safe
 *        to not initialize them yet.
 *      - The initialization functions always set strm->total_in and
 *        strm->total_out to zero.
 *      - If the initialization function fails, no memory is left allocated
 *        that would require freeing with lzma_end() even if some memory was
 *        associated with the lzma_stream structure when the initialization
 *        function was called.
 *
 *  - Use lzma_code() to do the actual work.
 *
 *  - Once the coding has been finished, the existing lzma_stream can be
 *    reused. It is OK to reuse lzma_stream with different initialization
 *    function without calling lzma_end() first. Old allocations are
 *    automatically freed.
 *
 *  - Finally, use lzma_end() to free the allocated memory. lzma_end() never
 *    frees the lzma_stream structure itself.
 *
 * Application may modify the values of total_in and total_out as it wants.
 * They are updated by liblzma to match the amount of data read and
 * written but aren't used for anything else except as a possible return
 * values from lzma_get_progress().
 */

type LzmaStream struct {
	NextIn        *uint8
	AvailIn       uintptr
	TotalIn       uint64
	NextOut       *uint8
	AvailOut      uintptr
	TotalOut      uint64
	Allocator     *LzmaAllocator
	Internal      *LzmaInternal
	ReservedPtr1  unsafe.Pointer
	ReservedPtr2  unsafe.Pointer
	ReservedPtr3  unsafe.Pointer
	ReservedPtr4  unsafe.Pointer
	SeekPos       uint64
	ReservedInt2  uint64
	ReservedInt3  uintptr
	ReservedInt4  uintptr
	ReservedEnum1 LzmaReservedEnum
	ReservedEnum2 LzmaReservedEnum
}

/**
 * \brief       Encode or decode data
 *
 * Once the lzma_stream has been successfully initialized (e.g. with
 * lzma_stream_encoder()), the actual encoding or decoding is done
 * using this function. The application has to update strm->next_in,
 * strm->avail_in, strm->next_out, and strm->avail_out to pass input
 * to and get output from liblzma.
 *
 * See the description of the coder-specific initialization function to find
 * out what `action' values are supported by the coder.
 *
 * \param       strm    Pointer to lzma_stream that is at least initialized
 *                      with LZMA_STREAM_INIT.
 * \param       action  Action for this function to take. Must be a valid
 *                      lzma_action enum value.
 *
 * \return      Any valid lzma_ret. See the lzma_ret enum description for more
 *              information.
 */
// llgo:link (*LzmaStream).LzmaCode C.lzma_code
func (recv_ *LzmaStream) LzmaCode(action LzmaAction) LzmaRet {
	return 0
}

/**
 * \brief       Free memory allocated for the coder data structures
 *
 * After lzma_end(strm), strm->internal is guaranteed to be NULL. No other
 * members of the lzma_stream structure are touched.
 *
 * \note        zlib indicates an error if application end()s unfinished
 *              stream structure. liblzma doesn't do this, and assumes that
 *              application knows what it is doing.
 *
 * \param       strm    Pointer to lzma_stream that is at least initialized
 *                      with LZMA_STREAM_INIT.
 */
// llgo:link (*LzmaStream).LzmaEnd C.lzma_end
func (recv_ *LzmaStream) LzmaEnd() {
}

/**
 * \brief       Get progress information
 *
 * In single-threaded mode, applications can get progress information from
 * strm->total_in and strm->total_out. In multi-threaded mode this is less
 * useful because a significant amount of both input and output data gets
 * buffered internally by liblzma. This makes total_in and total_out give
 * misleading information and also makes the progress indicator updates
 * non-smooth.
 *
 * This function gives realistic progress information also in multi-threaded
 * mode by taking into account the progress made by each thread. In
 * single-threaded mode *progress_in and *progress_out are set to
 * strm->total_in and strm->total_out, respectively.
 *
 * \param       strm          Pointer to lzma_stream that is at least
 *                            initialized with LZMA_STREAM_INIT.
 * \param[out]  progress_in   Pointer to the number of input bytes processed.
 * \param[out]  progress_out  Pointer to the number of output bytes processed.
 */
// llgo:link (*LzmaStream).LzmaGetProgress C.lzma_get_progress
func (recv_ *LzmaStream) LzmaGetProgress(progress_in *uint64, progress_out *uint64) {
}

/**
 * \brief       Get the memory usage of decoder filter chain
 *
 * This function is currently supported only when *strm has been initialized
 * with a function that takes a memlimit argument. With other functions, you
 * should use e.g. lzma_raw_encoder_memusage() or lzma_raw_decoder_memusage()
 * to estimate the memory requirements.
 *
 * This function is useful e.g. after LZMA_MEMLIMIT_ERROR to find out how big
 * the memory usage limit should have been to decode the input. Note that
 * this may give misleading information if decoding .xz Streams that have
 * multiple Blocks, because each Block can have different memory requirements.
 *
 * \param       strm    Pointer to lzma_stream that is at least initialized
 *                      with LZMA_STREAM_INIT.
 *
 * \return      How much memory is currently allocated for the filter
 *              decoders. If no filter chain is currently allocated,
 *              some non-zero value is still returned, which is less than
 *              or equal to what any filter chain would indicate as its
 *              memory requirement.
 *
 *              If this function isn't supported by *strm or some other error
 *              occurs, zero is returned.
 */
// llgo:link (*LzmaStream).LzmaMemusage C.lzma_memusage
func (recv_ *LzmaStream) LzmaMemusage() uint64 {
	return 0
}

/**
 * \brief       Get the current memory usage limit
 *
 * This function is supported only when *strm has been initialized with
 * a function that takes a memlimit argument.
 *
 * \param       strm    Pointer to lzma_stream that is at least initialized
 *                      with LZMA_STREAM_INIT.
 *
 * \return      On success, the current memory usage limit is returned
 *              (always non-zero). On error, zero is returned.
 */
// llgo:link (*LzmaStream).LzmaMemlimitGet C.lzma_memlimit_get
func (recv_ *LzmaStream) LzmaMemlimitGet() uint64 {
	return 0
}

/**
 * \brief       Set the memory usage limit
 *
 * This function is supported only when *strm has been initialized with
 * a function that takes a memlimit argument.
 *
 * liblzma 5.2.3 and earlier has a bug where memlimit value of 0 causes
 * this function to do nothing (leaving the limit unchanged) and still
 * return LZMA_OK. Later versions treat 0 as if 1 had been specified (so
 * lzma_memlimit_get() will return 1 even if you specify 0 here).
 *
 * liblzma 5.2.6 and earlier had a bug in single-threaded .xz decoder
 * (lzma_stream_decoder()) which made it impossible to continue decoding
 * after LZMA_MEMLIMIT_ERROR even if the limit was increased using
 * lzma_memlimit_set(). Other decoders worked correctly.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: New memory usage limit successfully set.
 *              - LZMA_MEMLIMIT_ERROR: The new limit is too small.
 *                The limit was not changed.
 *              - LZMA_PROG_ERROR: Invalid arguments, e.g. *strm doesn't
 *                support memory usage limit.
 */
// llgo:link (*LzmaStream).LzmaMemlimitSet C.lzma_memlimit_set
func (recv_ *LzmaStream) LzmaMemlimitSet(memlimit uint64) LzmaRet {
	return 0
}

type LzmaVli uint64

/**
 * \brief       Encode a variable-length integer
 *
 * This function has two modes: single-call and multi-call. Single-call mode
 * encodes the whole integer at once; it is an error if the output buffer is
 * too small. Multi-call mode saves the position in *vli_pos, and thus it is
 * possible to continue encoding if the buffer becomes full before the whole
 * integer has been encoded.
 *
 * \param       vli       Integer to be encoded
 * \param[out]  vli_pos   How many VLI-encoded bytes have already been written
 *                        out. When starting to encode a new integer in
 *                        multi-call mode, *vli_pos must be set to zero.
 *                        To use single-call encoding, set vli_pos to NULL.
 * \param[out]  out       Beginning of the output buffer
 * \param[out]  out_pos   The next byte will be written to out[*out_pos].
 * \param       out_size  Size of the out buffer; the first byte into
 *                        which no data is written to is out[out_size].
 *
 * \return      Slightly different return values are used in multi-call and
 *              single-call modes.
 *
 *              Single-call (vli_pos == NULL):
 *              - LZMA_OK: Integer successfully encoded.
 *              - LZMA_PROG_ERROR: Arguments are not sane. This can be due
 *                to too little output space; single-call mode doesn't use
 *                LZMA_BUF_ERROR, since the application should have checked
 *                the encoded size with lzma_vli_size().
 *
 *              Multi-call (vli_pos != NULL):
 *              - LZMA_OK: So far all OK, but the integer is not
 *                completely written out yet.
 *              - LZMA_STREAM_END: Integer successfully encoded.
 *              - LZMA_BUF_ERROR: No output space was provided.
 *              - LZMA_PROG_ERROR: Arguments are not sane.
 */
// llgo:link LzmaVli.LzmaVliEncode C.lzma_vli_encode
func (recv_ LzmaVli) LzmaVliEncode(vli_pos *uintptr, out *uint8, out_pos *uintptr, out_size uintptr) LzmaRet {
	return 0
}

/**
 * \brief       Decode a variable-length integer
 *
 * Like lzma_vli_encode(), this function has single-call and multi-call modes.
 *
 * \param[out]  vli       Pointer to decoded integer. The decoder will
 *                        initialize it to zero when *vli_pos == 0, so
 *                        application isn't required to initialize *vli.
 * \param[out]  vli_pos   How many bytes have already been decoded. When
 *                        starting to decode a new integer in multi-call
 *                        mode, *vli_pos must be initialized to zero. To
 *                        use single-call decoding, set vli_pos to NULL.
 * \param       in        Beginning of the input buffer
 * \param[out]  in_pos    The next byte will be read from in[*in_pos].
 * \param       in_size   Size of the input buffer; the first byte that
 *                        won't be read is in[in_size].
 *
 * \return      Slightly different return values are used in multi-call and
 *              single-call modes.
 *
 *              Single-call (vli_pos == NULL):
 *              - LZMA_OK: Integer successfully decoded.
 *              - LZMA_DATA_ERROR: Integer is corrupt. This includes hitting
 *                the end of the input buffer before the whole integer was
 *                decoded; providing no input at all will use LZMA_DATA_ERROR.
 *              - LZMA_PROG_ERROR: Arguments are not sane.
 *
 *              Multi-call (vli_pos != NULL):
 *              - LZMA_OK: So far all OK, but the integer is not
 *                completely decoded yet.
 *              - LZMA_STREAM_END: Integer successfully decoded.
 *              - LZMA_DATA_ERROR: Integer is corrupt.
 *              - LZMA_BUF_ERROR: No input was provided.
 *              - LZMA_PROG_ERROR: Arguments are not sane.
 */
// llgo:link (*LzmaVli).LzmaVliDecode C.lzma_vli_decode
func (recv_ *LzmaVli) LzmaVliDecode(vli_pos *uintptr, in *uint8, in_pos *uintptr, in_size uintptr) LzmaRet {
	return 0
}

/**
 * \brief       Get the number of bytes required to encode a VLI
 *
 * \param       vli       Integer whose encoded size is to be determined
 *
 * \return      Number of bytes on success (1-9). If vli isn't valid,
 *              zero is returned.
 */
// llgo:link LzmaVli.LzmaVliSize C.lzma_vli_size
func (recv_ LzmaVli) LzmaVliSize() uint32 {
	return 0
}

type LzmaCheck c.Int

const (
	LZMACHECKNONE   LzmaCheck = 0
	LZMACHECKCRC32  LzmaCheck = 1
	LZMACHECKCRC64  LzmaCheck = 4
	LZMACHECKSHA256 LzmaCheck = 10
)

/**
 * \brief       Test if the given Check ID is supported
 *
 * LZMA_CHECK_NONE and LZMA_CHECK_CRC32 are always supported (even if
 * liblzma is built with limited features).
 *
 * \note        It is safe to call this with a value that is not in the
 *              range [0, 15]; in that case the return value is always false.
 *
 * \param       check   Check ID
 *
 * \return      lzma_bool:
 *              - true if Check ID is supported by this liblzma build.
 *              - false otherwise.
 */
// llgo:link LzmaCheck.LzmaCheckIsSupported C.lzma_check_is_supported
func (recv_ LzmaCheck) LzmaCheckIsSupported() LzmaBool {
	return 0
}

/**
 * \brief       Get the size of the Check field with the given Check ID
 *
 * Although not all Check IDs have a check algorithm associated, the size of
 * every Check is already frozen. This function returns the size (in bytes) of
 * the Check field with the specified Check ID. The values are:
 * { 0, 4, 4, 4, 8, 8, 8, 16, 16, 16, 32, 32, 32, 64, 64, 64 }
 *
 * \param       check   Check ID
 *
 * \return      Size of the Check field in bytes. If the argument is not in
 *              the range [0, 15], UINT32_MAX is returned.
 */
// llgo:link LzmaCheck.LzmaCheckSize C.lzma_check_size
func (recv_ LzmaCheck) LzmaCheckSize() uint32 {
	return 0
}

/**
 * \brief       Calculate CRC32
 *
 * Calculate CRC32 using the polynomial from the IEEE 802.3 standard.
 *
 * \param       buf     Pointer to the input buffer
 * \param       size    Size of the input buffer
 * \param       crc     Previously returned CRC value. This is used to
 *                      calculate the CRC of a big buffer in smaller chunks.
 *                      Set to zero when starting a new calculation.
 *
 * \return      Updated CRC value, which can be passed to this function
 *              again to continue CRC calculation.
 */
//go:linkname LzmaCrc32 C.lzma_crc32
func LzmaCrc32(buf *uint8, size uintptr, crc uint32) uint32

/**
 * \brief       Calculate CRC64
 *
 * Calculate CRC64 using the polynomial from the ECMA-182 standard.
 *
 * This function is used similarly to lzma_crc32().
 *
 * \param       buf     Pointer to the input buffer
 * \param       size    Size of the input buffer
 * \param       crc     Previously returned CRC value. This is used to
 *                      calculate the CRC of a big buffer in smaller chunks.
 *                      Set to zero when starting a new calculation.
 *
 * \return      Updated CRC value, which can be passed to this function
 *              again to continue CRC calculation.
 */
//go:linkname LzmaCrc64 C.lzma_crc64
func LzmaCrc64(buf *uint8, size uintptr, crc uint64) uint64

/**
 * \brief       Get the type of the integrity check
 *
 * This function can be called only immediately after lzma_code() has
 * returned LZMA_NO_CHECK, LZMA_UNSUPPORTED_CHECK, or LZMA_GET_CHECK.
 * Calling this function in any other situation has undefined behavior.
 *
 * \param       strm    Pointer to lzma_stream meeting the above conditions.
 *
 * \return      Check ID in the lzma_stream, or undefined if called improperly.
 */
// llgo:link (*LzmaStream).LzmaGetCheck C.lzma_get_check
func (recv_ *LzmaStream) LzmaGetCheck() LzmaCheck {
	return 0
}

/**
 * \brief       Filter options
 *
 * This structure is used to pass a Filter ID and a pointer to the filter's
 * options to liblzma. A few functions work with a single lzma_filter
 * structure, while most functions expect a filter chain.
 *
 * A filter chain is indicated with an array of lzma_filter structures.
 * The array is terminated with .id = LZMA_VLI_UNKNOWN. Thus, the filter
 * array must have LZMA_FILTERS_MAX + 1 elements (that is, five) to
 * be able to hold any arbitrary filter chain. This is important when
 * using lzma_block_header_decode() from block.h, because a filter array
 * that is too small would make liblzma write past the end of the array.
 */

type LzmaFilter struct {
	Id      LzmaVli
	Options unsafe.Pointer
}

/**
* \brief       Test if the given Filter ID is supported for encoding
*
* \param       id      Filter ID
*
* \return      lzma_bool:
*              - true if the Filter ID is supported for encoding by this
*                liblzma build.
 *             - false otherwise.
*/
// llgo:link LzmaVli.LzmaFilterEncoderIsSupported C.lzma_filter_encoder_is_supported
func (recv_ LzmaVli) LzmaFilterEncoderIsSupported() LzmaBool {
	return 0
}

/**
 * \brief       Test if the given Filter ID is supported for decoding
 *
 * \param       id      Filter ID
 *
 * \return      lzma_bool:
 *              - true if the Filter ID is supported for decoding by this
 *                liblzma build.
 *              - false otherwise.
 */
// llgo:link LzmaVli.LzmaFilterDecoderIsSupported C.lzma_filter_decoder_is_supported
func (recv_ LzmaVli) LzmaFilterDecoderIsSupported() LzmaBool {
	return 0
}

/**
 * \brief       Copy the filters array
 *
 * Copy the Filter IDs and filter-specific options from src to dest.
 * Up to LZMA_FILTERS_MAX filters are copied, plus the terminating
 * .id == LZMA_VLI_UNKNOWN. Thus, dest should have at least
 * LZMA_FILTERS_MAX + 1 elements space unless the caller knows that
 * src is smaller than that.
 *
 * Unless the filter-specific options is NULL, the Filter ID has to be
 * supported by liblzma, because liblzma needs to know the size of every
 * filter-specific options structure. The filter-specific options are not
 * validated. If options is NULL, any unsupported Filter IDs are copied
 * without returning an error.
 *
 * Old filter-specific options in dest are not freed, so dest doesn't
 * need to be initialized by the caller in any way.
 *
 * If an error occurs, memory possibly already allocated by this function
 * is always freed. liblzma versions older than 5.2.7 may modify the dest
 * array and leave its contents in an undefined state if an error occurs.
 * liblzma 5.2.7 and newer only modify the dest array when returning LZMA_OK.
 *
 * \param       src         Array of filters terminated with
 *                          .id == LZMA_VLI_UNKNOWN.
 * \param[out]  dest        Destination filter array
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_MEM_ERROR
 *              - LZMA_OPTIONS_ERROR: Unsupported Filter ID and its options
 *                is not NULL.
 *              - LZMA_PROG_ERROR: src or dest is NULL.
 */
// llgo:link (*LzmaFilter).LzmaFiltersCopy C.lzma_filters_copy
func (recv_ *LzmaFilter) LzmaFiltersCopy(dest *LzmaFilter, allocator *LzmaAllocator) LzmaRet {
	return 0
}

/**
 * \brief       Free the options in the array of lzma_filter structures
 *
 * This frees the filter chain options. The filters array itself is not freed.
 *
 * The filters array must have at most LZMA_FILTERS_MAX + 1 elements
 * including the terminating element which must have .id = LZMA_VLI_UNKNOWN.
 * For all elements before the terminating element:
 *   - options will be freed using the given lzma_allocator or,
 *     if allocator is NULL, using free().
 *   - options will be set to NULL.
 *   - id will be set to LZMA_VLI_UNKNOWN.
 *
 * If filters is NULL, this does nothing. Again, this never frees the
 * filters array itself.
 *
 * \param       filters     Array of filters terminated with
 *                          .id == LZMA_VLI_UNKNOWN.
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 */
// llgo:link (*LzmaFilter).LzmaFiltersFree C.lzma_filters_free
func (recv_ *LzmaFilter) LzmaFiltersFree(allocator *LzmaAllocator) {
}

/**
 * \brief       Calculate approximate memory requirements for raw encoder
 *
 * This function can be used to calculate the memory requirements for
 * Block and Stream encoders too because Block and Stream encoders don't
 * need significantly more memory than raw encoder.
 *
 * \param       filters     Array of filters terminated with
 *                          .id == LZMA_VLI_UNKNOWN.
 *
 * \return      Number of bytes of memory required for the given
 *              filter chain when encoding or UINT64_MAX on error.
 */
// llgo:link (*LzmaFilter).LzmaRawEncoderMemusage C.lzma_raw_encoder_memusage
func (recv_ *LzmaFilter) LzmaRawEncoderMemusage() uint64 {
	return 0
}

/**
 * \brief       Calculate approximate memory requirements for raw decoder
 *
 * This function can be used to calculate the memory requirements for
 * Block and Stream decoders too because Block and Stream decoders don't
 * need significantly more memory than raw decoder.
 *
 * \param       filters     Array of filters terminated with
 *                          .id == LZMA_VLI_UNKNOWN.
 *
 * \return      Number of bytes of memory required for the given
 *              filter chain when decoding or UINT64_MAX on error.
 */
// llgo:link (*LzmaFilter).LzmaRawDecoderMemusage C.lzma_raw_decoder_memusage
func (recv_ *LzmaFilter) LzmaRawDecoderMemusage() uint64 {
	return 0
}

/**
 * \brief       Initialize raw encoder
 *
 * This function may be useful when implementing custom file formats.
 *
 * The `action' with lzma_code() can be LZMA_RUN, LZMA_SYNC_FLUSH (if the
 * filter chain supports it), or LZMA_FINISH.
 *
 * \param       strm      Pointer to lzma_stream that is at least
 *                        initialized with LZMA_STREAM_INIT.
 * \param       filters   Array of filters terminated with
 *                        .id == LZMA_VLI_UNKNOWN.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_MEM_ERROR
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaStream).LzmaRawEncoder C.lzma_raw_encoder
func (recv_ *LzmaStream) LzmaRawEncoder(filters *LzmaFilter) LzmaRet {
	return 0
}

/**
 * \brief       Initialize raw decoder
 *
 * The initialization of raw decoder goes similarly to raw encoder.
 *
 * The `action' with lzma_code() can be LZMA_RUN or LZMA_FINISH. Using
 * LZMA_FINISH is not required, it is supported just for convenience.
 *
 * \param       strm      Pointer to lzma_stream that is at least
 *                        initialized with LZMA_STREAM_INIT.
 * \param       filters   Array of filters terminated with
 *                        .id == LZMA_VLI_UNKNOWN.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_MEM_ERROR
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaStream).LzmaRawDecoder C.lzma_raw_decoder
func (recv_ *LzmaStream) LzmaRawDecoder(filters *LzmaFilter) LzmaRet {
	return 0
}

/**
 * \brief       Update the filter chain in the encoder
 *
 * This function may be called after lzma_code() has returned LZMA_STREAM_END
 * when LZMA_FULL_BARRIER, LZMA_FULL_FLUSH, or LZMA_SYNC_FLUSH was used:
 *
 *  - After LZMA_FULL_BARRIER or LZMA_FULL_FLUSH: Single-threaded .xz Stream
 *    encoder (lzma_stream_encoder()) and (since liblzma 5.4.0) multi-threaded
 *    Stream encoder (lzma_stream_encoder_mt()) allow setting a new filter
 *    chain to be used for the next Block(s).
 *
 *  - After LZMA_SYNC_FLUSH: Raw encoder (lzma_raw_encoder()),
 *    Block encoder (lzma_block_encoder()), and single-threaded .xz Stream
 *    encoder (lzma_stream_encoder()) allow changing certain filter-specific
 *    options in the middle of encoding. The actual filters in the chain
 *    (Filter IDs) must not be changed! Currently only the lc, lp, and pb
 *    options of LZMA2 (not LZMA1) can be changed this way.
 *
 *  - In the future some filters might allow changing some of their options
 *    without any barrier or flushing but currently such filters don't exist.
 *
 * This function may also be called when no data has been compressed yet
 * although this is rarely useful. In that case, this function will behave
 * as if LZMA_FULL_FLUSH (Stream encoders) or LZMA_SYNC_FLUSH (Raw or Block
 * encoder) had been used right before calling this function.
 *
 * \param       strm      Pointer to lzma_stream that is at least
 *                        initialized with LZMA_STREAM_INIT.
 * \param       filters   Array of filters terminated with
 *                        .id == LZMA_VLI_UNKNOWN.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_MEM_ERROR
 *              - LZMA_MEMLIMIT_ERROR
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaStream).LzmaFiltersUpdate C.lzma_filters_update
func (recv_ *LzmaStream) LzmaFiltersUpdate(filters *LzmaFilter) LzmaRet {
	return 0
}

/**
 * \brief       Single-call raw encoder
 *
 * \note        There is no function to calculate how big output buffer
 *              would surely be big enough. (lzma_stream_buffer_bound()
 *              works only for lzma_stream_buffer_encode(); raw encoder
 *              won't necessarily meet that bound.)
 *
 * \param       filters     Array of filters terminated with
 *                          .id == LZMA_VLI_UNKNOWN.
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 * \param       in          Beginning of the input buffer
 * \param       in_size     Size of the input buffer
 * \param[out]  out         Beginning of the output buffer
 * \param[out]  out_pos     The next byte will be written to out[*out_pos].
 *                          *out_pos is updated only if encoding succeeds.
 * \param       out_size    Size of the out buffer; the first byte into
 *                          which no data is written to is out[out_size].
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Encoding was successful.
 *              - LZMA_BUF_ERROR: Not enough output buffer space.
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_MEM_ERROR
 *              - LZMA_DATA_ERROR
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaFilter).LzmaRawBufferEncode C.lzma_raw_buffer_encode
func (recv_ *LzmaFilter) LzmaRawBufferEncode(allocator *LzmaAllocator, in *uint8, in_size uintptr, out *uint8, out_pos *uintptr, out_size uintptr) LzmaRet {
	return 0
}

/**
 * \brief       Single-call raw decoder
 *
 * \param       filters     Array of filters terminated with
 *                          .id == LZMA_VLI_UNKNOWN.
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 * \param       in          Beginning of the input buffer
 * \param       in_pos      The next byte will be read from in[*in_pos].
 *                          *in_pos is updated only if decoding succeeds.
 * \param       in_size     Size of the input buffer; the first byte that
 *                          won't be read is in[in_size].
 * \param[out]  out         Beginning of the output buffer
 * \param[out]  out_pos     The next byte will be written to out[*out_pos].
 *                          *out_pos is updated only if encoding succeeds.
 * \param       out_size    Size of the out buffer; the first byte into
 *                          which no data is written to is out[out_size].
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Decoding was successful.
 *              - LZMA_BUF_ERROR: Not enough output buffer space.
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_MEM_ERROR
 *              - LZMA_DATA_ERROR
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaFilter).LzmaRawBufferDecode C.lzma_raw_buffer_decode
func (recv_ *LzmaFilter) LzmaRawBufferDecode(allocator *LzmaAllocator, in *uint8, in_pos *uintptr, in_size uintptr, out *uint8, out_pos *uintptr, out_size uintptr) LzmaRet {
	return 0
}

/**
 * \brief       Get the size of the Filter Properties field
 *
 * This function may be useful when implementing custom file formats
 * using the raw encoder and decoder.
 *
 * \note        This function validates the Filter ID, but does not
 *              necessarily validate the options. Thus, it is possible
 *              that this returns LZMA_OK while the following call to
 *              lzma_properties_encode() returns LZMA_OPTIONS_ERROR.
 *
 * \param[out]  size    Pointer to uint32_t to hold the size of the properties
 * \param       filter  Filter ID and options (the size of the properties may
 *                      vary depending on the options)
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_PROG_ERROR
 */
//go:linkname LzmaPropertiesSize C.lzma_properties_size
func LzmaPropertiesSize(size *uint32, filter *LzmaFilter) LzmaRet

/**
 * \brief       Encode the Filter Properties field
 *
 * \note        Even this function won't validate more options than actually
 *              necessary. Thus, it is possible that encoding the properties
 *              succeeds but using the same options to initialize the encoder
 *              will fail.
 *
 * \note        If lzma_properties_size() indicated that the size
 *              of the Filter Properties field is zero, calling
 *              lzma_properties_encode() is not required, but it
 *              won't do any harm either.
 *
 * \param       filter  Filter ID and options
 * \param[out]  props   Buffer to hold the encoded options. The size of
 *                      the buffer must have been already determined with
 *                      lzma_properties_size().
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaFilter).LzmaPropertiesEncode C.lzma_properties_encode
func (recv_ *LzmaFilter) LzmaPropertiesEncode(props *uint8) LzmaRet {
	return 0
}

/**
 * \brief       Decode the Filter Properties field
 *
 * \param       filter      filter->id must have been set to the correct
 *                          Filter ID. filter->options doesn't need to be
 *                          initialized (it's not freed by this function). The
 *                          decoded options will be stored in filter->options;
 *                          it's application's responsibility to free it when
 *                          appropriate. filter->options is set to NULL if
 *                          there are no properties or if an error occurs.
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 *                          and in case of an error, also free().
 * \param       props       Input buffer containing the properties.
 * \param       props_size  Size of the properties. This must be the exact
 *                          size; giving too much or too little input will
 *                          return LZMA_OPTIONS_ERROR.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_MEM_ERROR
 */
// llgo:link (*LzmaFilter).LzmaPropertiesDecode C.lzma_properties_decode
func (recv_ *LzmaFilter) LzmaPropertiesDecode(allocator *LzmaAllocator, props *uint8, props_size uintptr) LzmaRet {
	return 0
}

/**
 * \brief       Calculate encoded size of a Filter Flags field
 *
 * Knowing the size of Filter Flags is useful to know when allocating
 * memory to hold the encoded Filter Flags.
 *
 * \note        If you need to calculate size of List of Filter Flags,
 *              you need to loop over every lzma_filter entry.
 *
 * \param[out]  size    Pointer to integer to hold the calculated size
 * \param       filter  Filter ID and associated options whose encoded
 *                      size is to be calculated
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: *size set successfully. Note that this doesn't
 *                guarantee that filter->options is valid, thus
 *                lzma_filter_flags_encode() may still fail.
 *              - LZMA_OPTIONS_ERROR: Unknown Filter ID or unsupported options.
 *              - LZMA_PROG_ERROR: Invalid options
 */
//go:linkname LzmaFilterFlagsSize C.lzma_filter_flags_size
func LzmaFilterFlagsSize(size *uint32, filter *LzmaFilter) LzmaRet

/**
 * \brief       Encode Filter Flags into given buffer
 *
 * In contrast to some functions, this doesn't allocate the needed buffer.
 * This is due to how this function is used internally by liblzma.
 *
 * \param       filter      Filter ID and options to be encoded
 * \param[out]  out         Beginning of the output buffer
 * \param[out]  out_pos     out[*out_pos] is the next write position. This
 *                          is updated by the encoder.
 * \param       out_size    out[out_size] is the first byte to not write.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Encoding was successful.
 *              - LZMA_OPTIONS_ERROR: Invalid or unsupported options.
 *              - LZMA_PROG_ERROR: Invalid options or not enough output
 *                buffer space (you should have checked it with
 *                lzma_filter_flags_size()).
 */
// llgo:link (*LzmaFilter).LzmaFilterFlagsEncode C.lzma_filter_flags_encode
func (recv_ *LzmaFilter) LzmaFilterFlagsEncode(out *uint8, out_pos *uintptr, out_size uintptr) LzmaRet {
	return 0
}

/**
 * \brief       Decode Filter Flags from given buffer
 *
 * The decoded result is stored into *filter. The old value of
 * filter->options is not free()d. If anything other than LZMA_OK
 * is returned, filter->options is set to NULL.
 *
 * \param[out]  filter      Destination filter. The decoded Filter ID will
 *                          be stored in filter->id. If options are needed
 *                          they will be allocated and the pointer will be
 *                          stored in filter->options.
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 * \param       in          Beginning of the input buffer
 * \param[out]  in_pos      The next byte will be read from in[*in_pos].
 *                          *in_pos is updated only if decoding succeeds.
 * \param       in_size     Size of the input buffer; the first byte that
 *                          won't be read is in[in_size].
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_MEM_ERROR
 *              - LZMA_DATA_ERROR
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaFilter).LzmaFilterFlagsDecode C.lzma_filter_flags_decode
func (recv_ *LzmaFilter) LzmaFilterFlagsDecode(allocator *LzmaAllocator, in *uint8, in_pos *uintptr, in_size uintptr) LzmaRet {
	return 0
}

/**
 * \brief       Convert a string to a filter chain
 *
 * This tries to make it easier to write applications that allow users
 * to set custom compression options. This only handles the filter
 * configuration (including presets) but not the number of threads,
 * block size, check type, or memory limits.
 *
 * The input string can be either a preset or a filter chain. Presets
 * begin with a digit 0-9 and may be followed by zero or more flags
 * which are lower-case letters. Currently only "e" is supported, matching
 * LZMA_PRESET_EXTREME. For partial xz command line syntax compatibility,
 * a preset string may start with a single dash "-".
 *
 * A filter chain consists of one or more "filtername:opt1=value1,opt2=value2"
 * strings separated by one or more spaces. Leading and trailing spaces are
 * ignored. All names and values must be lower-case. Extra commas in the
 * option list are ignored. The order of filters is significant: when
 * encoding, the uncompressed input data goes to the leftmost filter first.
 * Normally "lzma2" is the last filter in the chain.
 *
 * If one wishes to avoid spaces, for example, to avoid shell quoting,
 * it is possible to use two dashes "--" instead of spaces to separate
 * the filters.
 *
 * For xz command line compatibility, each filter may be prefixed with
 * two dashes "--" and the colon ":" separating the filter name from
 * the options may be replaced with an equals sign "=".
 *
 * By default, only filters that can be used in the .xz format are accepted.
 * To allow all filters (LZMA1) use the flag LZMA_STR_ALL_FILTERS.
 *
 * By default, very basic validation is done for the filter chain as a whole,
 * for example, that LZMA2 is only used as the last filter in the chain.
 * The validation isn't perfect though and it's possible that this function
 * succeeds but using the filter chain for encoding or decoding will still
 * result in LZMA_OPTIONS_ERROR. To disable this validation, use the flag
 * LZMA_STR_NO_VALIDATION.
 *
 * The available filter names and their options are available via
 * lzma_str_list_filters(). See the xz man page for the description
 * of filter names and options.
 *
 * For command line applications, below is an example how an error message
 * can be displayed. Note the use of an empty string for the field width.
 * If "^" was used there it would create an off-by-one error except at
 * the very beginning of the line.
 *
 * \code{.c}
 * const char *str = ...; // From user
 * lzma_filter filters[LZMA_FILTERS_MAX + 1];
 * int pos;
 * const char *msg = lzma_str_to_filters(str, &pos, filters, 0, NULL);
 * if (msg != NULL) {
 *     printf("%s: Error in XZ compression options:\n", argv[0]);
 *     printf("%s: %s\n", argv[0], str);
 *     printf("%s: %*s^\n", argv[0], errpos, "");
 *     printf("%s: %s\n", argv[0], msg);
 * }
 * \endcode
 *
 * \param       str         User-supplied string describing a preset or
 *                          a filter chain. If a default value is needed and
 *                          you don't know what would be good, use "6" since
 *                          that is the default preset in xz too.
 * \param[out]  error_pos   If this isn't NULL, this value will be set on
 *                          both success and on all errors. This tells the
 *                          location of the error in the string. This is
 *                          an int to make it straightforward to use this
 *                          as printf() field width. The value is guaranteed
 *                          to be in the range [0, INT_MAX] even if strlen(str)
 *                          somehow was greater than INT_MAX.
 * \param[out]  filters     An array of lzma_filter structures. There must
 *                          be LZMA_FILTERS_MAX + 1 (that is, five) elements
 *                          in the array. The old contents are ignored so it
 *                          doesn't need to be initialized. This array is
 *                          modified only if this function returns NULL.
 *                          Once the allocated filter options are no longer
 *                          needed, lzma_filters_free() can be used to free the
 *                          options (it doesn't free the filters array itself).
 * \param       flags       Bitwise-or of zero or more of the flags
 *                          LZMA_STR_ALL_FILTERS and LZMA_STR_NO_VALIDATION.
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 *
 * \return      On success, NULL is returned. On error, a statically-allocated
 *              error message is returned which together with the error_pos
 *              should give some idea what is wrong.
 */
//go:linkname LzmaStrToFilters C.lzma_str_to_filters
func LzmaStrToFilters(str *int8, error_pos *c.Int, filters *LzmaFilter, flags uint32, allocator *LzmaAllocator) *int8

/**
 * \brief       Convert a filter chain to a string
 *
 * Use cases:
 *
 *   - Verbose output showing the full encoder options to the user
 *     (use LZMA_STR_ENCODER in flags)
 *
 *   - Showing the filters and options that are required to decode a file
 *     (use LZMA_STR_DECODER in flags)
 *
 *   - Showing the filter names without any options in informational messages
 *     where the technical details aren't important (no flags). In this case
 *     the .options in the filters array are ignored and may be NULL even if
 *     a filter has a mandatory options structure.
 *
 * Note that even if the filter chain was specified using a preset,
 * the resulting filter chain isn't reversed to a preset. So if you
 * specify "6" to lzma_str_to_filters() then lzma_str_from_filters()
 * will produce a string containing "lzma2".
 *
 * \param[out]  str         On success *str will be set to point to an
 *                          allocated string describing the given filter
 *                          chain. Old value is ignored. On error *str is
 *                          always set to NULL.
 * \param       filters     Array of filters terminated with
 *                          .id == LZMA_VLI_UNKNOWN.
 * \param       flags       Bitwise-or of zero or more of the flags
 *                          LZMA_STR_ENCODER, LZMA_STR_DECODER,
 *                          LZMA_STR_GETOPT_LONG, and LZMA_STR_NO_SPACES.
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_OPTIONS_ERROR: Empty filter chain
 *                (filters[0].id == LZMA_VLI_UNKNOWN) or the filter chain
 *                includes a Filter ID that is not supported by this function.
 *              - LZMA_MEM_ERROR
 *              - LZMA_PROG_ERROR
 */
//go:linkname LzmaStrFromFilters C.lzma_str_from_filters
func LzmaStrFromFilters(str **int8, filters *LzmaFilter, flags uint32, allocator *LzmaAllocator) LzmaRet

/**
 * \brief       List available filters and/or their options (for help message)
 *
 * If a filter_id is given then only one line is created which contains the
 * filter name. If LZMA_STR_ENCODER or LZMA_STR_DECODER is used then the
 * options read by the encoder or decoder are printed on the same line.
 *
 * If filter_id is LZMA_VLI_UNKNOWN then all supported .xz-compatible filters
 * are listed:
 *
 *   - If neither LZMA_STR_ENCODER nor LZMA_STR_DECODER is used then
 *     the supported filter names are listed on a single line separated
 *     by spaces.
 *
 *   - If LZMA_STR_ENCODER or LZMA_STR_DECODER is used then filters and
 *     the supported options are listed one filter per line. There won't
 *     be a newline after the last filter.
 *
 *   - If LZMA_STR_ALL_FILTERS is used then the list will include also
 *     those filters that cannot be used in the .xz format (LZMA1).
 *
 * \param       str         On success *str will be set to point to an
 *                          allocated string listing the filters and options.
 *                          Old value is ignored. On error *str is always set
 *                          to NULL.
 * \param       filter_id   Filter ID or LZMA_VLI_UNKNOWN.
 * \param       flags       Bitwise-or of zero or more of the flags
 *                          LZMA_STR_ALL_FILTERS, LZMA_STR_ENCODER,
 *                          LZMA_STR_DECODER, and LZMA_STR_GETOPT_LONG.
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_OPTIONS_ERROR: Unsupported filter_id or flags
 *              - LZMA_MEM_ERROR
 *              - LZMA_PROG_ERROR
 */
//go:linkname LzmaStrListFilters C.lzma_str_list_filters
func LzmaStrListFilters(str **int8, filter_id LzmaVli, flags uint32, allocator *LzmaAllocator) LzmaRet

/**
 * \brief       Options for BCJ filters
 *
 * The BCJ filters never change the size of the data. Specifying options
 * for them is optional: if pointer to options is NULL, default value is
 * used. You probably never need to specify options to BCJ filters, so just
 * set the options pointer to NULL and be happy.
 *
 * If options with non-default values have been specified when encoding,
 * the same options must also be specified when decoding.
 *
 * \note        At the moment, none of the BCJ filters support
 *              LZMA_SYNC_FLUSH. If LZMA_SYNC_FLUSH is specified,
 *              LZMA_OPTIONS_ERROR will be returned. If there is need,
 *              partial support for LZMA_SYNC_FLUSH can be added in future.
 *              Partial means that flushing would be possible only at
 *              offsets that are multiple of 2, 4, or 16 depending on
 *              the filter, except x86 which cannot be made to support
 *              LZMA_SYNC_FLUSH predictably.
 */

type LzmaOptionsBcj struct {
	StartOffset uint32
}
type LzmaDeltaType c.Int

const LZMADELTATYPEBYTE LzmaDeltaType = 0

/**
 * \brief       Options for the Delta filter
 *
 * These options are needed by both encoder and decoder.
 */

type LzmaOptionsDelta struct {
	Type         LzmaDeltaType
	Dist         uint32
	ReservedInt1 uint32
	ReservedInt2 uint32
	ReservedInt3 uint32
	ReservedInt4 uint32
	ReservedPtr1 unsafe.Pointer
	ReservedPtr2 unsafe.Pointer
}
type LzmaMatchFinder c.Int

const (
	LZMAMFHC3 LzmaMatchFinder = 3
	LZMAMFHC4 LzmaMatchFinder = 4
	LZMAMFBT2 LzmaMatchFinder = 18
	LZMAMFBT3 LzmaMatchFinder = 19
	LZMAMFBT4 LzmaMatchFinder = 20
)

/**
 * \brief       Test if given match finder is supported
 *
 * It is safe to call this with a value that isn't listed in
 * lzma_match_finder enumeration; the return value will be false.
 *
 * There is no way to list which match finders are available in this
 * particular liblzma version and build. It would be useless, because
 * a new match finder, which the application developer wasn't aware,
 * could require giving additional options to the encoder that the older
 * match finders don't need.
 *
 * \param       match_finder    Match finder ID
 *
 * \return      lzma_bool:
 *              - true if the match finder is supported by this liblzma build.
 *              - false otherwise.
 */
// llgo:link LzmaMatchFinder.LzmaMfIsSupported C.lzma_mf_is_supported
func (recv_ LzmaMatchFinder) LzmaMfIsSupported() LzmaBool {
	return 0
}

type LzmaMode c.Int

const (
	LZMAMODEFAST   LzmaMode = 1
	LZMAMODENORMAL LzmaMode = 2
)

/**
 * \brief       Test if given compression mode is supported
 *
 * It is safe to call this with a value that isn't listed in lzma_mode
 * enumeration; the return value will be false.
 *
 * There is no way to list which modes are available in this particular
 * liblzma version and build. It would be useless, because a new compression
 * mode, which the application developer wasn't aware, could require giving
 * additional options to the encoder that the older modes don't need.
 *
 * \param       mode    Mode ID.
 *
 * \return      lzma_bool:
 *              - true if the compression mode is supported by this liblzma
 *                build.
 *              - false otherwise.
 */
// llgo:link LzmaMode.LzmaModeIsSupported C.lzma_mode_is_supported
func (recv_ LzmaMode) LzmaModeIsSupported() LzmaBool {
	return 0
}

/**
 * \brief       Options specific to the LZMA1 and LZMA2 filters
 *
 * Since LZMA1 and LZMA2 share most of the code, it's simplest to share
 * the options structure too. For encoding, all but the reserved variables
 * need to be initialized unless specifically mentioned otherwise.
 * lzma_lzma_preset() can be used to get a good starting point.
 *
 * For raw decoding, both LZMA1 and LZMA2 need dict_size, preset_dict, and
 * preset_dict_size (if preset_dict != NULL). LZMA1 needs also lc, lp, and pb.
 */

type LzmaOptionsLzma struct {
	DictSize       uint32
	PresetDict     *uint8
	PresetDictSize uint32
	Lc             uint32
	Lp             uint32
	Pb             uint32
	Mode           LzmaMode
	NiceLen        uint32
	Mf             LzmaMatchFinder
	Depth          uint32
	ExtFlags       uint32
	ExtSizeLow     uint32
	ExtSizeHigh    uint32
	ReservedInt4   uint32
	ReservedInt5   uint32
	ReservedInt6   uint32
	ReservedInt7   uint32
	ReservedInt8   uint32
	ReservedEnum1  LzmaReservedEnum
	ReservedEnum2  LzmaReservedEnum
	ReservedEnum3  LzmaReservedEnum
	ReservedEnum4  LzmaReservedEnum
	ReservedPtr1   unsafe.Pointer
	ReservedPtr2   unsafe.Pointer
}

/**
 * \brief       Set a compression preset to lzma_options_lzma structure
 *
 * 0 is the fastest and 9 is the slowest. These match the switches -0 .. -9
 * of the xz command line tool. In addition, it is possible to bitwise-or
 * flags to the preset. Currently only LZMA_PRESET_EXTREME is supported.
 * The flags are defined in container.h, because the flags are used also
 * with lzma_easy_encoder().
 *
 * The preset levels are subject to changes between liblzma versions.
 *
 * This function is available only if LZMA1 or LZMA2 encoder has been enabled
 * when building liblzma.
 *
 * If features (like certain match finders) have been disabled at build time,
 * then the function may return success (false) even though the resulting
 * LZMA1/LZMA2 options may not be usable for encoder initialization
 * (LZMA_OPTIONS_ERROR).
 *
 * \param[out]  options Pointer to LZMA1 or LZMA2 options to be filled
 * \param       preset  Preset level bitwse-ORed with preset flags
 *
 * \return      lzma_bool:
 *              - true if the preset is not supported (failure).
 *              - false otherwise (success).
 */
// llgo:link (*LzmaOptionsLzma).LzmaLzmaPreset C.lzma_lzma_preset
func (recv_ *LzmaOptionsLzma) LzmaLzmaPreset(preset uint32) LzmaBool {
	return 0
}

/**
 * \brief       Multithreading options
 */

type LzmaMt struct {
	Flags             uint32
	Threads           uint32
	BlockSize         uint64
	Timeout           uint32
	Preset            uint32
	Filters           *LzmaFilter
	Check             LzmaCheck
	ReservedEnum1     LzmaReservedEnum
	ReservedEnum2     LzmaReservedEnum
	ReservedEnum3     LzmaReservedEnum
	ReservedInt1      uint32
	ReservedInt2      uint32
	ReservedInt3      uint32
	ReservedInt4      uint32
	MemlimitThreading uint64
	MemlimitStop      uint64
	ReservedInt7      uint64
	ReservedInt8      uint64
	ReservedPtr1      unsafe.Pointer
	ReservedPtr2      unsafe.Pointer
	ReservedPtr3      unsafe.Pointer
	ReservedPtr4      unsafe.Pointer
}

/**
 * \brief       Calculate approximate memory usage of easy encoder
 *
 * This function is a wrapper for lzma_raw_encoder_memusage().
 *
 * \param       preset  Compression preset (level and possible flags)
 *
 * \return      Number of bytes of memory required for the given
 *              preset when encoding or UINT64_MAX on error.
 */
//go:linkname LzmaEasyEncoderMemusage C.lzma_easy_encoder_memusage
func LzmaEasyEncoderMemusage(preset uint32) uint64

/**
 * \brief       Calculate approximate decoder memory usage of a preset
 *
 * This function is a wrapper for lzma_raw_decoder_memusage().
 *
 * \param       preset  Compression preset (level and possible flags)
 *
 * \return      Number of bytes of memory required to decompress a file
 *              that was compressed using the given preset or UINT64_MAX
 *              on error.
 */
//go:linkname LzmaEasyDecoderMemusage C.lzma_easy_decoder_memusage
func LzmaEasyDecoderMemusage(preset uint32) uint64

/**
 * \brief       Initialize .xz Stream encoder using a preset number
 *
 * This function is intended for those who just want to use the basic features
 * of liblzma (that is, most developers out there).
 *
 * If initialization fails (return value is not LZMA_OK), all the memory
 * allocated for *strm by liblzma is always freed. Thus, there is no need
 * to call lzma_end() after failed initialization.
 *
 * If initialization succeeds, use lzma_code() to do the actual encoding.
 * Valid values for `action' (the second argument of lzma_code()) are
 * LZMA_RUN, LZMA_SYNC_FLUSH, LZMA_FULL_FLUSH, and LZMA_FINISH. In future,
 * there may be compression levels or flags that don't support LZMA_SYNC_FLUSH.
 *
 * \param       strm    Pointer to lzma_stream that is at least initialized
 *                      with LZMA_STREAM_INIT.
 * \param       preset  Compression preset to use. A preset consist of level
 *                      number and zero or more flags. Usually flags aren't
 *                      used, so preset is simply a number [0, 9] which match
 *                      the options -0 ... -9 of the xz command line tool.
 *                      Additional flags can be be set using bitwise-or with
 *                      the preset level number, e.g. 6 | LZMA_PRESET_EXTREME.
 * \param       check   Integrity check type to use. See check.h for available
 *                      checks. The xz command line tool defaults to
 *                      LZMA_CHECK_CRC64, which is a good choice if you are
 *                      unsure. LZMA_CHECK_CRC32 is good too as long as the
 *                      uncompressed file is not many gigabytes.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Initialization succeeded. Use lzma_code() to
 *                encode your data.
 *              - LZMA_MEM_ERROR: Memory allocation failed.
 *              - LZMA_OPTIONS_ERROR: The given compression preset is not
 *                supported by this build of liblzma.
 *              - LZMA_UNSUPPORTED_CHECK: The given check type is not
 *                supported by this liblzma build.
 *              - LZMA_PROG_ERROR: One or more of the parameters have values
 *                that will never be valid. For example, strm == NULL.
 */
// llgo:link (*LzmaStream).LzmaEasyEncoder C.lzma_easy_encoder
func (recv_ *LzmaStream) LzmaEasyEncoder(preset uint32, check LzmaCheck) LzmaRet {
	return 0
}

/**
 * \brief       Single-call .xz Stream encoding using a preset number
 *
 * The maximum required output buffer size can be calculated with
 * lzma_stream_buffer_bound().
 *
 * \param       preset      Compression preset to use. See the description
 *                          in lzma_easy_encoder().
 * \param       check       Type of the integrity check to calculate from
 *                          uncompressed data.
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 * \param       in          Beginning of the input buffer
 * \param       in_size     Size of the input buffer
 * \param[out]  out         Beginning of the output buffer
 * \param[out]  out_pos     The next byte will be written to out[*out_pos].
 *                          *out_pos is updated only if encoding succeeds.
 * \param       out_size    Size of the out buffer; the first byte into
 *                          which no data is written to is out[out_size].
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Encoding was successful.
 *              - LZMA_BUF_ERROR: Not enough output buffer space.
 *              - LZMA_UNSUPPORTED_CHECK
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_MEM_ERROR
 *              - LZMA_DATA_ERROR
 *              - LZMA_PROG_ERROR
 */
//go:linkname LzmaEasyBufferEncode C.lzma_easy_buffer_encode
func LzmaEasyBufferEncode(preset uint32, check LzmaCheck, allocator *LzmaAllocator, in *uint8, in_size uintptr, out *uint8, out_pos *uintptr, out_size uintptr) LzmaRet

/**
 * \brief       Initialize .xz Stream encoder using a custom filter chain
 *
 * \param       strm    Pointer to lzma_stream that is at least initialized
 *                      with LZMA_STREAM_INIT.
 * \param       filters Array of filters terminated with
 *                      .id == LZMA_VLI_UNKNOWN. See filters.h for more
 *                      information.
 * \param       check   Type of the integrity check to calculate from
 *                      uncompressed data.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Initialization was successful.
 *              - LZMA_MEM_ERROR
 *              - LZMA_UNSUPPORTED_CHECK
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaStream).LzmaStreamEncoder C.lzma_stream_encoder
func (recv_ *LzmaStream) LzmaStreamEncoder(filters *LzmaFilter, check LzmaCheck) LzmaRet {
	return 0
}

/**
 * \brief       Calculate approximate memory usage of multithreaded .xz encoder
 *
 * Since doing the encoding in threaded mode doesn't affect the memory
 * requirements of single-threaded decompressor, you can use
 * lzma_easy_decoder_memusage(options->preset) or
 * lzma_raw_decoder_memusage(options->filters) to calculate
 * the decompressor memory requirements.
 *
 * \param       options Compression options
 *
 * \return      Number of bytes of memory required for encoding with the
 *              given options. If an error occurs, for example due to
 *              unsupported preset or filter chain, UINT64_MAX is returned.
 */
// llgo:link (*LzmaMt).LzmaStreamEncoderMtMemusage C.lzma_stream_encoder_mt_memusage
func (recv_ *LzmaMt) LzmaStreamEncoderMtMemusage() uint64 {
	return 0
}

/**
 * \brief       Initialize multithreaded .xz Stream encoder
 *
 * This provides the functionality of lzma_easy_encoder() and
 * lzma_stream_encoder() as a single function for multithreaded use.
 *
 * The supported actions for lzma_code() are LZMA_RUN, LZMA_FULL_FLUSH,
 * LZMA_FULL_BARRIER, and LZMA_FINISH. Support for LZMA_SYNC_FLUSH might be
 * added in the future.
 *
 * \param       strm    Pointer to lzma_stream that is at least initialized
 *                      with LZMA_STREAM_INIT.
 * \param       options Pointer to multithreaded compression options
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_MEM_ERROR
 *              - LZMA_UNSUPPORTED_CHECK
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaStream).LzmaStreamEncoderMt C.lzma_stream_encoder_mt
func (recv_ *LzmaStream) LzmaStreamEncoderMt(options *LzmaMt) LzmaRet {
	return 0
}

/**
 * \brief       Initialize .lzma encoder (legacy file format)
 *
 * The .lzma format is sometimes called the LZMA_Alone format, which is the
 * reason for the name of this function. The .lzma format supports only the
 * LZMA1 filter. There is no support for integrity checks like CRC32.
 *
 * Use this function if and only if you need to create files readable by
 * legacy LZMA tools such as LZMA Utils 4.32.x. Moving to the .xz format
 * is strongly recommended.
 *
 * The valid action values for lzma_code() are LZMA_RUN and LZMA_FINISH.
 * No kind of flushing is supported, because the file format doesn't make
 * it possible.
 *
 * \param       strm    Pointer to lzma_stream that is at least initialized
 *                      with LZMA_STREAM_INIT.
 * \param       options Pointer to encoder options
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_MEM_ERROR
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaStream).LzmaAloneEncoder C.lzma_alone_encoder
func (recv_ *LzmaStream) LzmaAloneEncoder(options *LzmaOptionsLzma) LzmaRet {
	return 0
}

/**
 * \brief       Calculate output buffer size for single-call Stream encoder
 *
 * When trying to compress incompressible data, the encoded size will be
 * slightly bigger than the input data. This function calculates how much
 * output buffer space is required to be sure that lzma_stream_buffer_encode()
 * doesn't return LZMA_BUF_ERROR.
 *
 * The calculated value is not exact, but it is guaranteed to be big enough.
 * The actual maximum output space required may be slightly smaller (up to
 * about 100 bytes). This should not be a problem in practice.
 *
 * If the calculated maximum size doesn't fit into size_t or would make the
 * Stream grow past LZMA_VLI_MAX (which should never happen in practice),
 * zero is returned to indicate the error.
 *
 * \note        The limit calculated by this function applies only to
 *              single-call encoding. Multi-call encoding may (and probably
 *              will) have larger maximum expansion when encoding
 *              incompressible data. Currently there is no function to
 *              calculate the maximum expansion of multi-call encoding.
 *
 * \param       uncompressed_size   Size in bytes of the uncompressed
 *                                  input data
 *
 * \return      Maximum number of bytes needed to store the compressed data.
 */
//go:linkname LzmaStreamBufferBound C.lzma_stream_buffer_bound
func LzmaStreamBufferBound(uncompressed_size uintptr) uintptr

/**
 * \brief       Single-call .xz Stream encoder
 *
 * \param       filters     Array of filters terminated with
 *                          .id == LZMA_VLI_UNKNOWN. See filters.h for more
 *                          information.
 * \param       check       Type of the integrity check to calculate from
 *                          uncompressed data.
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 * \param       in          Beginning of the input buffer
 * \param       in_size     Size of the input buffer
 * \param[out]  out         Beginning of the output buffer
 * \param[out]  out_pos     The next byte will be written to out[*out_pos].
 *                          *out_pos is updated only if encoding succeeds.
 * \param       out_size    Size of the out buffer; the first byte into
 *                          which no data is written to is out[out_size].
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Encoding was successful.
 *              - LZMA_BUF_ERROR: Not enough output buffer space.
 *              - LZMA_UNSUPPORTED_CHECK
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_MEM_ERROR
 *              - LZMA_DATA_ERROR
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaFilter).LzmaStreamBufferEncode C.lzma_stream_buffer_encode
func (recv_ *LzmaFilter) LzmaStreamBufferEncode(check LzmaCheck, allocator *LzmaAllocator, in *uint8, in_size uintptr, out *uint8, out_pos *uintptr, out_size uintptr) LzmaRet {
	return 0
}

/**
 * \brief       MicroLZMA encoder
 *
 * The MicroLZMA format is a raw LZMA stream whose first byte (always 0x00)
 * has been replaced with bitwise-negation of the LZMA properties (lc/lp/pb).
 * This encoding ensures that the first byte of MicroLZMA stream is never
 * 0x00. There is no end of payload marker and thus the uncompressed size
 * must be stored separately. For the best error detection the dictionary
 * size should be stored separately as well but alternatively one may use
 * the uncompressed size as the dictionary size when decoding.
 *
 * With the MicroLZMA encoder, lzma_code() behaves slightly unusually.
 * The action argument must be LZMA_FINISH and the return value will never be
 * LZMA_OK. Thus the encoding is always done with a single lzma_code() after
 * the initialization. The benefit of the combination of initialization
 * function and lzma_code() is that memory allocations can be re-used for
 * better performance.
 *
 * lzma_code() will try to encode as much input as is possible to fit into
 * the given output buffer. If not all input can be encoded, the stream will
 * be finished without encoding all the input. The caller must check both
 * input and output buffer usage after lzma_code() (total_in and total_out
 * in lzma_stream can be convenient). Often lzma_code() can fill the output
 * buffer completely if there is a lot of input, but sometimes a few bytes
 * may remain unused because the next LZMA symbol would require more space.
 *
 * lzma_stream.avail_out must be at least 6. Otherwise LZMA_PROG_ERROR
 * will be returned.
 *
 * The LZMA dictionary should be reasonably low to speed up the encoder
 * re-initialization. A good value is bigger than the resulting
 * uncompressed size of most of the output chunks. For example, if output
 * size is 4 KiB, dictionary size of 32 KiB or 64 KiB is good. If the
 * data compresses extremely well, even 128 KiB may be useful.
 *
 * The MicroLZMA format and this encoder variant were made with the EROFS
 * file system in mind. This format may be convenient in other embedded
 * uses too where many small streams are needed. XZ Embedded includes a
 * decoder for this format.
 *
 * \param       strm    Pointer to lzma_stream that is at least initialized
 *                      with LZMA_STREAM_INIT.
 * \param       options Pointer to encoder options
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_STREAM_END: All good. Check the amounts of input used
 *                and output produced. Store the amount of input used
 *                (uncompressed size) as it needs to be known to decompress
 *                the data.
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_MEM_ERROR
 *              - LZMA_PROG_ERROR: In addition to the generic reasons for this
 *                error code, this may also be returned if there isn't enough
 *                output space (6 bytes) to create a valid MicroLZMA stream.
 */
// llgo:link (*LzmaStream).LzmaMicrolzmaEncoder C.lzma_microlzma_encoder
func (recv_ *LzmaStream) LzmaMicrolzmaEncoder(options *LzmaOptionsLzma) LzmaRet {
	return 0
}

/**
 * \brief       Initialize .xz Stream decoder
 *
 * \param       strm        Pointer to lzma_stream that is at least initialized
 *                          with LZMA_STREAM_INIT.
 * \param       memlimit    Memory usage limit as bytes. Use UINT64_MAX
 *                          to effectively disable the limiter. liblzma
 *                          5.2.3 and earlier don't allow 0 here and return
 *                          LZMA_PROG_ERROR; later versions treat 0 as if 1
 *                          had been specified.
 * \param       flags       Bitwise-or of zero or more of the decoder flags:
 *                          LZMA_TELL_NO_CHECK, LZMA_TELL_UNSUPPORTED_CHECK,
 *                          LZMA_TELL_ANY_CHECK, LZMA_IGNORE_CHECK,
 *                          LZMA_CONCATENATED, LZMA_FAIL_FAST
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Initialization was successful.
 *              - LZMA_MEM_ERROR: Cannot allocate memory.
 *              - LZMA_OPTIONS_ERROR: Unsupported flags
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaStream).LzmaStreamDecoder C.lzma_stream_decoder
func (recv_ *LzmaStream) LzmaStreamDecoder(memlimit uint64, flags uint32) LzmaRet {
	return 0
}

/**
 * \brief       Initialize multithreaded .xz Stream decoder
 *
 * The decoder can decode multiple Blocks in parallel. This requires that each
 * Block Header contains the Compressed Size and Uncompressed size fields
 * which are added by the multi-threaded encoder, see lzma_stream_encoder_mt().
 *
 * A Stream with one Block will only utilize one thread. A Stream with multiple
 * Blocks but without size information in Block Headers will be processed in
 * single-threaded mode in the same way as done by lzma_stream_decoder().
 * Concatenated Streams are processed one Stream at a time; no inter-Stream
 * parallelization is done.
 *
 * This function behaves like lzma_stream_decoder() when options->threads == 1
 * and options->memlimit_threading <= 1.
 *
 * \param       strm        Pointer to lzma_stream that is at least initialized
 *                          with LZMA_STREAM_INIT.
 * \param       options     Pointer to multithreaded compression options
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Initialization was successful.
 *              - LZMA_MEM_ERROR: Cannot allocate memory.
 *              - LZMA_MEMLIMIT_ERROR: Memory usage limit was reached.
 *              - LZMA_OPTIONS_ERROR: Unsupported flags.
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaStream).LzmaStreamDecoderMt C.lzma_stream_decoder_mt
func (recv_ *LzmaStream) LzmaStreamDecoderMt(options *LzmaMt) LzmaRet {
	return 0
}

/**
* \brief       Decode .xz, .lzma, and .lz (lzip) files with autodetection
*
* This decoder autodetects between the .xz, .lzma, and .lz file formats,
* and calls lzma_stream_decoder(), lzma_alone_decoder(), or
* lzma_lzip_decoder() once the type of the input file has been detected.
*
* Support for .lz was added in 5.4.0.
*
* If the flag LZMA_CONCATENATED is used and the input is a .lzma file:
* For historical reasons concatenated .lzma files aren't supported.
* If there is trailing data after one .lzma stream, lzma_code() will
* return LZMA_DATA_ERROR. (lzma_alone_decoder() doesn't have such a check
* as it doesn't support any decoder flags. It will return LZMA_STREAM_END
* after one .lzma stream.)
*
 * \param       strm       Pointer to lzma_stream that is at least initialized
*                          with LZMA_STREAM_INIT.
* \param       memlimit    Memory usage limit as bytes. Use UINT64_MAX
*                          to effectively disable the limiter. liblzma
*                          5.2.3 and earlier don't allow 0 here and return
*                          LZMA_PROG_ERROR; later versions treat 0 as if 1
*                          had been specified.
* \param       flags       Bitwise-or of zero or more of the decoder flags:
*                          LZMA_TELL_NO_CHECK, LZMA_TELL_UNSUPPORTED_CHECK,
*                          LZMA_TELL_ANY_CHECK, LZMA_IGNORE_CHECK,
*                          LZMA_CONCATENATED, LZMA_FAIL_FAST
*
* \return      Possible lzma_ret values:
*              - LZMA_OK: Initialization was successful.
*              - LZMA_MEM_ERROR: Cannot allocate memory.
*              - LZMA_OPTIONS_ERROR: Unsupported flags
*              - LZMA_PROG_ERROR
*/
// llgo:link (*LzmaStream).LzmaAutoDecoder C.lzma_auto_decoder
func (recv_ *LzmaStream) LzmaAutoDecoder(memlimit uint64, flags uint32) LzmaRet {
	return 0
}

/**
 * \brief       Initialize .lzma decoder (legacy file format)
 *
 * Valid `action' arguments to lzma_code() are LZMA_RUN and LZMA_FINISH.
 * There is no need to use LZMA_FINISH, but it's allowed because it may
 * simplify certain types of applications.
 *
 * \param       strm        Pointer to lzma_stream that is at least initialized
 *                          with LZMA_STREAM_INIT.
 * \param       memlimit    Memory usage limit as bytes. Use UINT64_MAX
 *                          to effectively disable the limiter. liblzma
 *                          5.2.3 and earlier don't allow 0 here and return
 *                          LZMA_PROG_ERROR; later versions treat 0 as if 1
 *                          had been specified.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_MEM_ERROR
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaStream).LzmaAloneDecoder C.lzma_alone_decoder
func (recv_ *LzmaStream) LzmaAloneDecoder(memlimit uint64) LzmaRet {
	return 0
}

/**
 * \brief       Initialize .lz (lzip) decoder (a foreign file format)
 *
 * This decoder supports the .lz format version 0 and the unextended .lz
 * format version 1:
 *
 *   - Files in the format version 0 were produced by lzip 1.3 and older.
 *     Such files aren't common but may be found from file archives
 *     as a few source packages were released in this format. People
 *     might have old personal files in this format too. Decompression
 *     support for the format version 0 was removed in lzip 1.18.
 *
 *   - lzip 1.3 added decompression support for .lz format version 1 files.
 *     Compression support was added in lzip 1.4. In lzip 1.6 the .lz format
 *     version 1 was extended to support the Sync Flush marker. This extension
 *     is not supported by liblzma. lzma_code() will return LZMA_DATA_ERROR
 *     at the location of the Sync Flush marker. In practice files with
 *     the Sync Flush marker are very rare and thus liblzma can decompress
 *     almost all .lz files.
 *
 * Just like with lzma_stream_decoder() for .xz files, LZMA_CONCATENATED
 * should be used when decompressing normal standalone .lz files.
 *
 * The .lz format allows putting non-.lz data at the end of a file after at
 * least one valid .lz member. That is, one can append custom data at the end
 * of a .lz file and the decoder is required to ignore it. In liblzma this
 * is relevant only when LZMA_CONCATENATED is used. In that case lzma_code()
 * will return LZMA_STREAM_END and leave lzma_stream.next_in pointing to
 * the first byte of the non-.lz data. An exception to this is if the first
 * 1-3 bytes of the non-.lz data are identical to the .lz magic bytes
 * (0x4C, 0x5A, 0x49, 0x50; "LZIP" in US-ASCII). In such a case the 1-3 bytes
 * will have been ignored by lzma_code(). If one wishes to locate the non-.lz
 * data reliably, one must ensure that the first byte isn't 0x4C. Actually
 * one should ensure that none of the first four bytes of trailing data are
 * equal to the magic bytes because lzip >= 1.20 requires it by default.
 *
 * \param       strm        Pointer to lzma_stream that is at least initialized
 *                          with LZMA_STREAM_INIT.
 * \param       memlimit    Memory usage limit as bytes. Use UINT64_MAX
 *                          to effectively disable the limiter.
 * \param       flags       Bitwise-or of flags, or zero for no flags.
 *                          All decoder flags listed above are supported
 *                          although only LZMA_CONCATENATED and (in very rare
 *                          cases) LZMA_IGNORE_CHECK are actually useful.
 *                          LZMA_TELL_NO_CHECK, LZMA_TELL_UNSUPPORTED_CHECK,
 *                          and LZMA_FAIL_FAST do nothing. LZMA_TELL_ANY_CHECK
 *                          is supported for consistency only as CRC32 is
 *                          always used in the .lz format.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Initialization was successful.
 *              - LZMA_MEM_ERROR: Cannot allocate memory.
 *              - LZMA_OPTIONS_ERROR: Unsupported flags
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaStream).LzmaLzipDecoder C.lzma_lzip_decoder
func (recv_ *LzmaStream) LzmaLzipDecoder(memlimit uint64, flags uint32) LzmaRet {
	return 0
}

/**
 * \brief       Single-call .xz Stream decoder
 *
 * \param       memlimit    Pointer to how much memory the decoder is allowed
 *                          to allocate. The value pointed by this pointer is
 *                          modified if and only if LZMA_MEMLIMIT_ERROR is
 *                          returned.
 * \param       flags       Bitwise-or of zero or more of the decoder flags:
 *                          LZMA_TELL_NO_CHECK, LZMA_TELL_UNSUPPORTED_CHECK,
 *                          LZMA_IGNORE_CHECK, LZMA_CONCATENATED,
 *                          LZMA_FAIL_FAST. Note that LZMA_TELL_ANY_CHECK
 *                          is not allowed and will return LZMA_PROG_ERROR.
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 * \param       in          Beginning of the input buffer
 * \param       in_pos      The next byte will be read from in[*in_pos].
 *                          *in_pos is updated only if decoding succeeds.
 * \param       in_size     Size of the input buffer; the first byte that
 *                          won't be read is in[in_size].
 * \param[out]  out         Beginning of the output buffer
 * \param[out]  out_pos     The next byte will be written to out[*out_pos].
 *                          *out_pos is updated only if decoding succeeds.
 * \param       out_size    Size of the out buffer; the first byte into
 *                          which no data is written to is out[out_size].
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Decoding was successful.
 *              - LZMA_FORMAT_ERROR
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_DATA_ERROR
 *              - LZMA_NO_CHECK: This can be returned only if using
 *                the LZMA_TELL_NO_CHECK flag.
 *              - LZMA_UNSUPPORTED_CHECK: This can be returned only if using
 *                the LZMA_TELL_UNSUPPORTED_CHECK flag.
 *              - LZMA_MEM_ERROR
 *              - LZMA_MEMLIMIT_ERROR: Memory usage limit was reached.
 *                The minimum required memlimit value was stored to *memlimit.
 *              - LZMA_BUF_ERROR: Output buffer was too small.
 *              - LZMA_PROG_ERROR
 */
//go:linkname LzmaStreamBufferDecode C.lzma_stream_buffer_decode
func LzmaStreamBufferDecode(memlimit *uint64, flags uint32, allocator *LzmaAllocator, in *uint8, in_pos *uintptr, in_size uintptr, out *uint8, out_pos *uintptr, out_size uintptr) LzmaRet

/**
 * \brief       MicroLZMA decoder
 *
 * See lzma_microlzma_encoder() for more information.
 *
 * The lzma_code() usage with this decoder is completely normal. The
 * special behavior of lzma_code() applies to lzma_microlzma_encoder() only.
 *
 * \param       strm        Pointer to lzma_stream that is at least initialized
 *                          with LZMA_STREAM_INIT.
 * \param       comp_size   Compressed size of the MicroLZMA stream.
 *                          The caller must somehow know this exactly.
 * \param       uncomp_size Uncompressed size of the MicroLZMA stream.
 *                          If the exact uncompressed size isn't known, this
 *                          can be set to a value that is at most as big as
 *                          the exact uncompressed size would be, but then the
 *                          next argument uncomp_size_is_exact must be false.
 * \param       uncomp_size_is_exact
 *                          If true, uncomp_size must be exactly correct.
 *                          This will improve error detection at the end of
 *                          the stream. If the exact uncompressed size isn't
 *                          known, this must be false. uncomp_size must still
 *                          be at most as big as the exact uncompressed size
 *                          is. Setting this to false when the exact size is
 *                          known will work but error detection at the end of
 *                          the stream will be weaker.
 * \param       dict_size   LZMA dictionary size that was used when
 *                          compressing the data. It is OK to use a bigger
 *                          value too but liblzma will then allocate more
 *                          memory than would actually be required and error
 *                          detection will be slightly worse. (Note that with
 *                          the implementation in XZ Embedded it doesn't
 *                          affect the memory usage if one specifies bigger
 *                          dictionary than actually required.)
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_MEM_ERROR
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaStream).LzmaMicrolzmaDecoder C.lzma_microlzma_decoder
func (recv_ *LzmaStream) LzmaMicrolzmaDecoder(comp_size uint64, uncomp_size uint64, uncomp_size_is_exact LzmaBool, dict_size uint32) LzmaRet {
	return 0
}

/**
 * \brief       Options for encoding/decoding Stream Header and Stream Footer
 */

type LzmaStreamFlags struct {
	Version       uint32
	BackwardSize  LzmaVli
	Check         LzmaCheck
	ReservedEnum1 LzmaReservedEnum
	ReservedEnum2 LzmaReservedEnum
	ReservedEnum3 LzmaReservedEnum
	ReservedEnum4 LzmaReservedEnum
	ReservedBool1 LzmaBool
	ReservedBool2 LzmaBool
	ReservedBool3 LzmaBool
	ReservedBool4 LzmaBool
	ReservedBool5 LzmaBool
	ReservedBool6 LzmaBool
	ReservedBool7 LzmaBool
	ReservedBool8 LzmaBool
	ReservedInt1  uint32
	ReservedInt2  uint32
}

/**
 * \brief       Encode Stream Header
 *
 * \param       options     Stream Header options to be encoded.
 *                          options->backward_size is ignored and doesn't
 *                          need to be initialized.
 * \param[out]  out         Beginning of the output buffer of
 *                          LZMA_STREAM_HEADER_SIZE bytes.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Encoding was successful.
 *              - LZMA_OPTIONS_ERROR: options->version is not supported by
 *                this liblzma version.
 *              - LZMA_PROG_ERROR: Invalid options.
 */
// llgo:link (*LzmaStreamFlags).LzmaStreamHeaderEncode C.lzma_stream_header_encode
func (recv_ *LzmaStreamFlags) LzmaStreamHeaderEncode(out *uint8) LzmaRet {
	return 0
}

/**
 * \brief       Encode Stream Footer
 *
 * \param       options     Stream Footer options to be encoded.
 * \param[out]  out         Beginning of the output buffer of
 *                          LZMA_STREAM_HEADER_SIZE bytes.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Encoding was successful.
 *              - LZMA_OPTIONS_ERROR: options->version is not supported by
 *                this liblzma version.
 *              - LZMA_PROG_ERROR: Invalid options.
 */
// llgo:link (*LzmaStreamFlags).LzmaStreamFooterEncode C.lzma_stream_footer_encode
func (recv_ *LzmaStreamFlags) LzmaStreamFooterEncode(out *uint8) LzmaRet {
	return 0
}

/**
 * \brief       Decode Stream Header
 *
 * options->backward_size is always set to LZMA_VLI_UNKNOWN. This is to
 * help comparing Stream Flags from Stream Header and Stream Footer with
 * lzma_stream_flags_compare().
 *
 * \note        When decoding .xz files that contain multiple Streams, it may
 *              make sense to print "file format not recognized" only if
 *              decoding of the Stream Header of the \a first Stream gives
 *              LZMA_FORMAT_ERROR. If non-first Stream Header gives
 *              LZMA_FORMAT_ERROR, the message used for LZMA_DATA_ERROR is
 *              probably more appropriate.
 *              For example, the Stream decoder in liblzma uses
 *              LZMA_DATA_ERROR if LZMA_FORMAT_ERROR is returned by
 *              lzma_stream_header_decode() when decoding non-first Stream.
 *
 * \param[out]  options     Target for the decoded Stream Header options.
 * \param       in          Beginning of the input buffer of
 *                          LZMA_STREAM_HEADER_SIZE bytes.
 *
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Decoding was successful.
 *              - LZMA_FORMAT_ERROR: Magic bytes don't match, thus the given
 *                buffer cannot be Stream Header.
 *              - LZMA_DATA_ERROR: CRC32 doesn't match, thus the header
 *                is corrupt.
 *              - LZMA_OPTIONS_ERROR: Unsupported options are present
 *                in the header.
 */
// llgo:link (*LzmaStreamFlags).LzmaStreamHeaderDecode C.lzma_stream_header_decode
func (recv_ *LzmaStreamFlags) LzmaStreamHeaderDecode(in *uint8) LzmaRet {
	return 0
}

/**
 * \brief       Decode Stream Footer
 *
 * \note        If Stream Header was already decoded successfully, but
 *              decoding Stream Footer returns LZMA_FORMAT_ERROR, the
 *              application should probably report some other error message
 *              than "file format not recognized". The file likely
 *              is corrupt (possibly truncated). The Stream decoder in liblzma
 *              uses LZMA_DATA_ERROR in this situation.
 *
 * \param[out]  options     Target for the decoded Stream Footer options.
 * \param       in          Beginning of the input buffer of
 *                          LZMA_STREAM_HEADER_SIZE bytes.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Decoding was successful.
 *              - LZMA_FORMAT_ERROR: Magic bytes don't match, thus the given
 *                buffer cannot be Stream Footer.
 *              - LZMA_DATA_ERROR: CRC32 doesn't match, thus the Stream Footer
 *                is corrupt.
 *              - LZMA_OPTIONS_ERROR: Unsupported options are present
 *                in Stream Footer.
 */
// llgo:link (*LzmaStreamFlags).LzmaStreamFooterDecode C.lzma_stream_footer_decode
func (recv_ *LzmaStreamFlags) LzmaStreamFooterDecode(in *uint8) LzmaRet {
	return 0
}

/**
 * \brief       Compare two lzma_stream_flags structures
 *
 * backward_size values are compared only if both are not
 * LZMA_VLI_UNKNOWN.
 *
 * \param       a       Pointer to lzma_stream_flags structure
 * \param       b       Pointer to lzma_stream_flags structure
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Both are equal. If either had backward_size set
 *                to LZMA_VLI_UNKNOWN, backward_size values were not
 *                compared or validated.
 *              - LZMA_DATA_ERROR: The structures differ.
 *              - LZMA_OPTIONS_ERROR: version in either structure is greater
 *                than the maximum supported version (currently zero).
 *              - LZMA_PROG_ERROR: Invalid value, e.g. invalid check or
 *                backward_size.
 */
// llgo:link (*LzmaStreamFlags).LzmaStreamFlagsCompare C.lzma_stream_flags_compare
func (recv_ *LzmaStreamFlags) LzmaStreamFlagsCompare(b *LzmaStreamFlags) LzmaRet {
	return 0
}

/**
 * \brief       Options for the Block and Block Header encoders and decoders
 *
 * Different Block handling functions use different parts of this structure.
 * Some read some members, other functions write, and some do both. Only the
 * members listed for reading need to be initialized when the specified
 * functions are called. The members marked for writing will be assigned
 * new values at some point either by calling the given function or by
 * later calls to lzma_code().
 */

type LzmaBlock struct {
	Version          uint32
	HeaderSize       uint32
	Check            LzmaCheck
	CompressedSize   LzmaVli
	UncompressedSize LzmaVli
	Filters          *LzmaFilter
	RawCheck         [64]uint8
	ReservedPtr1     unsafe.Pointer
	ReservedPtr2     unsafe.Pointer
	ReservedPtr3     unsafe.Pointer
	ReservedInt1     uint32
	ReservedInt2     uint32
	ReservedInt3     LzmaVli
	ReservedInt4     LzmaVli
	ReservedInt5     LzmaVli
	ReservedInt6     LzmaVli
	ReservedInt7     LzmaVli
	ReservedInt8     LzmaVli
	ReservedEnum1    LzmaReservedEnum
	ReservedEnum2    LzmaReservedEnum
	ReservedEnum3    LzmaReservedEnum
	ReservedEnum4    LzmaReservedEnum
	IgnoreCheck      LzmaBool
	ReservedBool2    LzmaBool
	ReservedBool3    LzmaBool
	ReservedBool4    LzmaBool
	ReservedBool5    LzmaBool
	ReservedBool6    LzmaBool
	ReservedBool7    LzmaBool
	ReservedBool8    LzmaBool
}

/**
 * \brief       Calculate Block Header Size
 *
 * Calculate the minimum size needed for the Block Header field using the
 * settings specified in the lzma_block structure. Note that it is OK to
 * increase the calculated header_size value as long as it is a multiple of
 * four and doesn't exceed LZMA_BLOCK_HEADER_SIZE_MAX. Increasing header_size
 * just means that lzma_block_header_encode() will add Header Padding.
 *
 * \note        This doesn't check that all the options are valid i.e. this
 *              may return LZMA_OK even if lzma_block_header_encode() or
 *              lzma_block_encoder() would fail. If you want to validate the
 *              filter chain, consider using lzma_memlimit_encoder() which as
 *              a side-effect validates the filter chain.
 *
 * \param       block   Block options
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Size calculated successfully and stored to
 *                block->header_size.
 *              - LZMA_OPTIONS_ERROR: Unsupported version, filters or
 *                filter options.
 *              - LZMA_PROG_ERROR: Invalid values like compressed_size == 0.
 */
// llgo:link (*LzmaBlock).LzmaBlockHeaderSize C.lzma_block_header_size
func (recv_ *LzmaBlock) LzmaBlockHeaderSize() LzmaRet {
	return 0
}

/**
 * \brief       Encode Block Header
 *
 * The caller must have calculated the size of the Block Header already with
 * lzma_block_header_size(). If a value larger than the one calculated by
 * lzma_block_header_size() is used, the Block Header will be padded to the
 * specified size.
 *
 * \param       block       Block options to be encoded.
 * \param[out]  out         Beginning of the output buffer. This must be
 *                          at least block->header_size bytes.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Encoding was successful. block->header_size
 *                bytes were written to output buffer.
 *              - LZMA_OPTIONS_ERROR: Invalid or unsupported options.
 *              - LZMA_PROG_ERROR: Invalid arguments, for example
 *                block->header_size is invalid or block->filters is NULL.
 */
// llgo:link (*LzmaBlock).LzmaBlockHeaderEncode C.lzma_block_header_encode
func (recv_ *LzmaBlock) LzmaBlockHeaderEncode(out *uint8) LzmaRet {
	return 0
}

/**
 * \brief       Decode Block Header
 *
 * block->version should (usually) be set to the highest value supported
 * by the application. If the application sets block->version to a value
 * higher than supported by the current liblzma version, this function will
 * downgrade block->version to the highest value supported by it. Thus one
 * should check the value of block->version after calling this function if
 * block->version was set to a non-zero value and the application doesn't
 * otherwise know that the liblzma version being used is new enough to
 * support the specified block->version.
 *
 * The size of the Block Header must have already been decoded with
 * lzma_block_header_size_decode() macro and stored to block->header_size.
 *
 * The integrity check type from Stream Header must have been stored
 * to block->check.
 *
 * block->filters must have been allocated, but they don't need to be
 * initialized (possible existing filter options are not freed).
 *
 * \param[out]  block       Destination for Block options
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() (and also free()
 *                          if an error occurs).
 * \param       in          Beginning of the input buffer. This must be
 *                          at least block->header_size bytes.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Decoding was successful. block->header_size
 *                bytes were read from the input buffer.
 *              - LZMA_OPTIONS_ERROR: The Block Header specifies some
 *                unsupported options such as unsupported filters. This can
 *                happen also if block->version was set to a too low value
 *                compared to what would be required to properly represent
 *                the information stored in the Block Header.
 *              - LZMA_DATA_ERROR: Block Header is corrupt, for example,
 *                the CRC32 doesn't match.
 *              - LZMA_PROG_ERROR: Invalid arguments, for example
 *                block->header_size is invalid or block->filters is NULL.
 */
// llgo:link (*LzmaBlock).LzmaBlockHeaderDecode C.lzma_block_header_decode
func (recv_ *LzmaBlock) LzmaBlockHeaderDecode(allocator *LzmaAllocator, in *uint8) LzmaRet {
	return 0
}

/**
 * \brief       Validate and set Compressed Size according to Unpadded Size
 *
 * Block Header stores Compressed Size, but Index has Unpadded Size. If the
 * application has already parsed the Index and is now decoding Blocks,
 * it can calculate Compressed Size from Unpadded Size. This function does
 * exactly that with error checking:
 *
 *  - Compressed Size calculated from Unpadded Size must be positive integer,
 *    that is, Unpadded Size must be big enough that after Block Header and
 *    Check fields there's still at least one byte for Compressed Size.
 *
 *  - If Compressed Size was present in Block Header, the new value
 *    calculated from Unpadded Size is compared against the value
 *    from Block Header.
 *
 * \note        This function must be called _after_ decoding the Block Header
 *              field so that it can properly validate Compressed Size if it
 *              was present in Block Header.
 *
 * \param       block           Block options: block->header_size must
 *                              already be set with lzma_block_header_size().
 * \param       unpadded_size   Unpadded Size from the Index field in bytes
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: block->compressed_size was set successfully.
 *              - LZMA_DATA_ERROR: unpadded_size is too small compared to
 *                block->header_size and lzma_check_size(block->check).
 *              - LZMA_PROG_ERROR: Some values are invalid. For example,
 *                block->header_size must be a multiple of four and
 *                between 8 and 1024 inclusive.
 */
// llgo:link (*LzmaBlock).LzmaBlockCompressedSize C.lzma_block_compressed_size
func (recv_ *LzmaBlock) LzmaBlockCompressedSize(unpadded_size LzmaVli) LzmaRet {
	return 0
}

/**
 * \brief       Calculate Unpadded Size
 *
 * The Index field stores Unpadded Size and Uncompressed Size. The latter
 * can be taken directly from the lzma_block structure after coding a Block,
 * but Unpadded Size needs to be calculated from Block Header Size,
 * Compressed Size, and size of the Check field. This is where this function
 * is needed.
 *
 * \param       block   Block options: block->header_size must already be
 *                      set with lzma_block_header_size().
 *
 * \return      Unpadded Size on success, or zero on error.
 */
// llgo:link (*LzmaBlock).LzmaBlockUnpaddedSize C.lzma_block_unpadded_size
func (recv_ *LzmaBlock) LzmaBlockUnpaddedSize() LzmaVli {
	return 0
}

/**
 * \brief       Calculate the total encoded size of a Block
 *
 * This is equivalent to lzma_block_unpadded_size() except that the returned
 * value includes the size of the Block Padding field.
 *
 * \param       block   Block options: block->header_size must already be
 *                      set with lzma_block_header_size().
 *
 * \return      On success, total encoded size of the Block. On error,
 *              zero is returned.
 */
// llgo:link (*LzmaBlock).LzmaBlockTotalSize C.lzma_block_total_size
func (recv_ *LzmaBlock) LzmaBlockTotalSize() LzmaVli {
	return 0
}

/**
 * \brief       Initialize .xz Block encoder
 *
 * Valid actions for lzma_code() are LZMA_RUN, LZMA_SYNC_FLUSH (only if the
 * filter chain supports it), and LZMA_FINISH.
 *
 * The Block encoder encodes the Block Data, Block Padding, and Check value.
 * It does NOT encode the Block Header which can be encoded with
 * lzma_block_header_encode().
 *
 * \param       strm    Pointer to lzma_stream that is at least initialized
 *                      with LZMA_STREAM_INIT.
 * \param       block   Block options: block->version, block->check,
 *                      and block->filters must have been initialized.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: All good, continue with lzma_code().
 *              - LZMA_MEM_ERROR
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_UNSUPPORTED_CHECK: block->check specifies a Check ID
 *                that is not supported by this build of liblzma. Initializing
 *                the encoder failed.
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaStream).LzmaBlockEncoder C.lzma_block_encoder
func (recv_ *LzmaStream) LzmaBlockEncoder(block *LzmaBlock) LzmaRet {
	return 0
}

/**
 * \brief       Initialize .xz Block decoder
 *
 * Valid actions for lzma_code() are LZMA_RUN and LZMA_FINISH. Using
 * LZMA_FINISH is not required. It is supported only for convenience.
 *
 * The Block decoder decodes the Block Data, Block Padding, and Check value.
 * It does NOT decode the Block Header which can be decoded with
 * lzma_block_header_decode().
 *
 * \param       strm    Pointer to lzma_stream that is at least initialized
 *                      with LZMA_STREAM_INIT.
 * \param       block   Block options
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: All good, continue with lzma_code().
 *              - LZMA_PROG_ERROR
 *              - LZMA_MEM_ERROR
 */
// llgo:link (*LzmaStream).LzmaBlockDecoder C.lzma_block_decoder
func (recv_ *LzmaStream) LzmaBlockDecoder(block *LzmaBlock) LzmaRet {
	return 0
}

/**
 * \brief       Calculate maximum output size for single-call Block encoding
 *
 * This is equivalent to lzma_stream_buffer_bound() but for .xz Blocks.
 * See the documentation of lzma_stream_buffer_bound().
 *
 * \param       uncompressed_size   Size of the data to be encoded with the
 *                                  single-call Block encoder.
 *
 * \return      Maximum output size in bytes for single-call Block encoding.
 */
//go:linkname LzmaBlockBufferBound C.lzma_block_buffer_bound
func LzmaBlockBufferBound(uncompressed_size uintptr) uintptr

/**
 * \brief       Single-call .xz Block encoder
 *
 * In contrast to the multi-call encoder initialized with
 * lzma_block_encoder(), this function encodes also the Block Header. This
 * is required to make it possible to write appropriate Block Header also
 * in case the data isn't compressible, and different filter chain has to be
 * used to encode the data in uncompressed form using uncompressed chunks
 * of the LZMA2 filter.
 *
 * When the data isn't compressible, header_size, compressed_size, and
 * uncompressed_size are set just like when the data was compressible, but
 * it is possible that header_size is too small to hold the filter chain
 * specified in block->filters, because that isn't necessarily the filter
 * chain that was actually used to encode the data. lzma_block_unpadded_size()
 * still works normally, because it doesn't read the filters array.
 *
 * \param       block       Block options: block->version, block->check,
 *                          and block->filters must have been initialized.
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 * \param       in          Beginning of the input buffer
 * \param       in_size     Size of the input buffer
 * \param[out]  out         Beginning of the output buffer
 * \param[out]  out_pos     The next byte will be written to out[*out_pos].
 *                          *out_pos is updated only if encoding succeeds.
 * \param       out_size    Size of the out buffer; the first byte into
 *                          which no data is written to is out[out_size].
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Encoding was successful.
 *              - LZMA_BUF_ERROR: Not enough output buffer space.
 *              - LZMA_UNSUPPORTED_CHECK
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_MEM_ERROR
 *              - LZMA_DATA_ERROR
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaBlock).LzmaBlockBufferEncode C.lzma_block_buffer_encode
func (recv_ *LzmaBlock) LzmaBlockBufferEncode(allocator *LzmaAllocator, in *uint8, in_size uintptr, out *uint8, out_pos *uintptr, out_size uintptr) LzmaRet {
	return 0
}

/**
 * \brief       Single-call uncompressed .xz Block encoder
 *
 * This is like lzma_block_buffer_encode() except this doesn't try to
 * compress the data and instead encodes the data using LZMA2 uncompressed
 * chunks. The required output buffer size can be determined with
 * lzma_block_buffer_bound().
 *
 * Since the data won't be compressed, this function ignores block->filters.
 * This function doesn't take lzma_allocator because this function doesn't
 * allocate any memory from the heap.
 *
 * \param       block       Block options: block->version, block->check,
 *                          and block->filters must have been initialized.
 * \param       in          Beginning of the input buffer
 * \param       in_size     Size of the input buffer
 * \param[out]  out         Beginning of the output buffer
 * \param[out]  out_pos     The next byte will be written to out[*out_pos].
 *                          *out_pos is updated only if encoding succeeds.
 * \param       out_size    Size of the out buffer; the first byte into
 *                          which no data is written to is out[out_size].
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Encoding was successful.
 *              - LZMA_BUF_ERROR: Not enough output buffer space.
 *              - LZMA_UNSUPPORTED_CHECK
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_MEM_ERROR
 *              - LZMA_DATA_ERROR
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaBlock).LzmaBlockUncompEncode C.lzma_block_uncomp_encode
func (recv_ *LzmaBlock) LzmaBlockUncompEncode(in *uint8, in_size uintptr, out *uint8, out_pos *uintptr, out_size uintptr) LzmaRet {
	return 0
}

/**
 * \brief       Single-call .xz Block decoder
 *
 * This is single-call equivalent of lzma_block_decoder(), and requires that
 * the caller has already decoded Block Header and checked its memory usage.
 *
 * \param       block       Block options
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 * \param       in          Beginning of the input buffer
 * \param       in_pos      The next byte will be read from in[*in_pos].
 *                          *in_pos is updated only if decoding succeeds.
 * \param       in_size     Size of the input buffer; the first byte that
 *                          won't be read is in[in_size].
 * \param[out]  out         Beginning of the output buffer
 * \param[out]  out_pos     The next byte will be written to out[*out_pos].
 *                          *out_pos is updated only if encoding succeeds.
 * \param       out_size    Size of the out buffer; the first byte into
 *                          which no data is written to is out[out_size].
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Decoding was successful.
 *              - LZMA_OPTIONS_ERROR
 *              - LZMA_DATA_ERROR
 *              - LZMA_MEM_ERROR
 *              - LZMA_BUF_ERROR: Output buffer was too small.
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaBlock).LzmaBlockBufferDecode C.lzma_block_buffer_decode
func (recv_ *LzmaBlock) LzmaBlockBufferDecode(allocator *LzmaAllocator, in *uint8, in_pos *uintptr, in_size uintptr, out *uint8, out_pos *uintptr, out_size uintptr) LzmaRet {
	return 0
}

type LzmaIndexS struct {
	Unused [8]uint8
}
type LzmaIndex LzmaIndexS

/**
 * \brief       Iterator to get information about Blocks and Streams
 */

type LzmaIndexIter struct {
	Stream struct {
		Flags              *LzmaStreamFlags
		ReservedPtr1       unsafe.Pointer
		ReservedPtr2       unsafe.Pointer
		ReservedPtr3       unsafe.Pointer
		Number             LzmaVli
		BlockCount         LzmaVli
		CompressedOffset   LzmaVli
		UncompressedOffset LzmaVli
		CompressedSize     LzmaVli
		UncompressedSize   LzmaVli
		Padding            LzmaVli
		ReservedVli1       LzmaVli
		ReservedVli2       LzmaVli
		ReservedVli3       LzmaVli
		ReservedVli4       LzmaVli
	}
	Block struct {
		NumberInFile             LzmaVli
		CompressedFileOffset     LzmaVli
		UncompressedFileOffset   LzmaVli
		NumberInStream           LzmaVli
		CompressedStreamOffset   LzmaVli
		UncompressedStreamOffset LzmaVli
		UncompressedSize         LzmaVli
		UnpaddedSize             LzmaVli
		TotalSize                LzmaVli
		ReservedVli1             LzmaVli
		ReservedVli2             LzmaVli
		ReservedVli3             LzmaVli
		ReservedVli4             LzmaVli
		ReservedPtr1             unsafe.Pointer
		ReservedPtr2             unsafe.Pointer
		ReservedPtr3             unsafe.Pointer
		ReservedPtr4             unsafe.Pointer
	}
	Internal [6]struct {
		P unsafe.Pointer
	}
}
type LzmaIndexIterMode c.Int

const (
	LZMAINDEXITERANY           LzmaIndexIterMode = 0
	LZMAINDEXITERSTREAM        LzmaIndexIterMode = 1
	LZMAINDEXITERBLOCK         LzmaIndexIterMode = 2
	LZMAINDEXITERNONEMPTYBLOCK LzmaIndexIterMode = 3
)

/**
 * \brief       Calculate memory usage of lzma_index
 *
 * On disk, the size of the Index field depends on both the number of Records
 * stored and the size of the Records (due to variable-length integer
 * encoding). When the Index is kept in lzma_index structure, the memory usage
 * depends only on the number of Records/Blocks stored in the Index(es), and
 * in case of concatenated lzma_indexes, the number of Streams. The size in
 * RAM is almost always significantly bigger than in the encoded form on disk.
 *
 * This function calculates an approximate amount of memory needed to hold
 * the given number of Streams and Blocks in lzma_index structure. This
 * value may vary between CPU architectures and also between liblzma versions
 * if the internal implementation is modified.
 *
 * \param       streams Number of Streams
 * \param       blocks  Number of Blocks
 *
 * \return      Approximate memory in bytes needed in a lzma_index structure.
 */
// llgo:link LzmaVli.LzmaIndexMemusage C.lzma_index_memusage
func (recv_ LzmaVli) LzmaIndexMemusage(blocks LzmaVli) uint64 {
	return 0
}

/**
 * \brief       Calculate the memory usage of an existing lzma_index
 *
 * This is a shorthand for lzma_index_memusage(lzma_index_stream_count(i),
 * lzma_index_block_count(i)).
 *
 * \param       i   Pointer to lzma_index structure
 *
 * \return      Approximate memory in bytes used by the lzma_index structure.
 */
// llgo:link (*LzmaIndex).LzmaIndexMemused C.lzma_index_memused
func (recv_ *LzmaIndex) LzmaIndexMemused() uint64 {
	return 0
}

/**
 * \brief       Allocate and initialize a new lzma_index structure
 *
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 *
 * \return      On success, a pointer to an empty initialized lzma_index is
 *              returned. If allocation fails, NULL is returned.
 */
// llgo:link (*LzmaAllocator).LzmaIndexInit C.lzma_index_init
func (recv_ *LzmaAllocator) LzmaIndexInit() *LzmaIndex {
	return nil
}

/**
 * \brief       Deallocate lzma_index
 *
 * If i is NULL, this does nothing.
 *
 * \param       i           Pointer to lzma_index structure to deallocate
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 */
// llgo:link (*LzmaIndex).LzmaIndexEnd C.lzma_index_end
func (recv_ *LzmaIndex) LzmaIndexEnd(allocator *LzmaAllocator) {
}

/**
 * \brief       Add a new Block to lzma_index
 *
 * \param       i                 Pointer to a lzma_index structure
 * \param       allocator         lzma_allocator for custom allocator
 *                                functions. Set to NULL to use malloc()
 *                                and free().
 * \param       unpadded_size     Unpadded Size of a Block. This can be
 *                                calculated with lzma_block_unpadded_size()
 *                                after encoding or decoding the Block.
 * \param       uncompressed_size Uncompressed Size of a Block. This can be
 *                                taken directly from lzma_block structure
 *                                after encoding or decoding the Block.
 *
 * Appending a new Block does not invalidate iterators. For example,
 * if an iterator was pointing to the end of the lzma_index, after
 * lzma_index_append() it is possible to read the next Block with
 * an existing iterator.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_MEM_ERROR
 *              - LZMA_DATA_ERROR: Compressed or uncompressed size of the
 *                Stream or size of the Index field would grow too big.
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaIndex).LzmaIndexAppend C.lzma_index_append
func (recv_ *LzmaIndex) LzmaIndexAppend(allocator *LzmaAllocator, unpadded_size LzmaVli, uncompressed_size LzmaVli) LzmaRet {
	return 0
}

/**
 * \brief       Set the Stream Flags
 *
 * Set the Stream Flags of the last (and typically the only) Stream
 * in lzma_index. This can be useful when reading information from the
 * lzma_index, because to decode Blocks, knowing the integrity check type
 * is needed.
 *
 * \param       i              Pointer to lzma_index structure
 * \param       stream_flags   Pointer to lzma_stream_flags structure. This
 *                             is copied into the internal preallocated
 *                             structure, so the caller doesn't need to keep
 *                             the flags' data available after calling this
 *                             function.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_OPTIONS_ERROR: Unsupported stream_flags->version.
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaIndex).LzmaIndexStreamFlags C.lzma_index_stream_flags
func (recv_ *LzmaIndex) LzmaIndexStreamFlags(stream_flags *LzmaStreamFlags) LzmaRet {
	return 0
}

/**
 * \brief       Get the types of integrity Checks
 *
 * If lzma_index_stream_flags() is used to set the Stream Flags for
 * every Stream, lzma_index_checks() can be used to get a bitmask to
 * indicate which Check types have been used. It can be useful e.g. if
 * showing the Check types to the user.
 *
 * The bitmask is 1 << check_id, e.g. CRC32 is 1 << 1 and SHA-256 is 1 << 10.
 *
 * \param       i   Pointer to lzma_index structure
 *
 * \return      Bitmask indicating which Check types are used in the lzma_index
 */
// llgo:link (*LzmaIndex).LzmaIndexChecks C.lzma_index_checks
func (recv_ *LzmaIndex) LzmaIndexChecks() uint32 {
	return 0
}

/**
 * \brief       Set the amount of Stream Padding
 *
 * Set the amount of Stream Padding of the last (and typically the only)
 * Stream in the lzma_index. This is needed when planning to do random-access
 * reading within multiple concatenated Streams.
 *
 * By default, the amount of Stream Padding is assumed to be zero bytes.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_DATA_ERROR: The file size would grow too big.
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaIndex).LzmaIndexStreamPadding C.lzma_index_stream_padding
func (recv_ *LzmaIndex) LzmaIndexStreamPadding(stream_padding LzmaVli) LzmaRet {
	return 0
}

/**
 * \brief       Get the number of Streams
 *
 * \param       i   Pointer to lzma_index structure
 *
 * \return      Number of Streams in the lzma_index
 */
// llgo:link (*LzmaIndex).LzmaIndexStreamCount C.lzma_index_stream_count
func (recv_ *LzmaIndex) LzmaIndexStreamCount() LzmaVli {
	return 0
}

/**
 * \brief       Get the number of Blocks
 *
 * This returns the total number of Blocks in lzma_index. To get number
 * of Blocks in individual Streams, use lzma_index_iter.
 *
 * \param       i   Pointer to lzma_index structure
 *
 * \return      Number of blocks in the lzma_index
 */
// llgo:link (*LzmaIndex).LzmaIndexBlockCount C.lzma_index_block_count
func (recv_ *LzmaIndex) LzmaIndexBlockCount() LzmaVli {
	return 0
}

/**
 * \brief       Get the size of the Index field as bytes
 *
 * This is needed to verify the Backward Size field in the Stream Footer.
 *
 * \param       i   Pointer to lzma_index structure
 *
 * \return      Size in bytes of the Index
 */
// llgo:link (*LzmaIndex).LzmaIndexSize C.lzma_index_size
func (recv_ *LzmaIndex) LzmaIndexSize() LzmaVli {
	return 0
}

/**
 * \brief       Get the total size of the Stream
 *
 * If multiple lzma_indexes have been combined, this works as if the Blocks
 * were in a single Stream. This is useful if you are going to combine
 * Blocks from multiple Streams into a single new Stream.
 *
 * \param       i   Pointer to lzma_index structure
 *
 * \return      Size in bytes of the Stream (if all Blocks are combined
 *              into one Stream).
 */
// llgo:link (*LzmaIndex).LzmaIndexStreamSize C.lzma_index_stream_size
func (recv_ *LzmaIndex) LzmaIndexStreamSize() LzmaVli {
	return 0
}

/**
 * \brief       Get the total size of the Blocks
 *
 * This doesn't include the Stream Header, Stream Footer, Stream Padding,
 * or Index fields.
 *
 * \param       i   Pointer to lzma_index structure
 *
 * \return      Size in bytes of all Blocks in the Stream(s)
 */
// llgo:link (*LzmaIndex).LzmaIndexTotalSize C.lzma_index_total_size
func (recv_ *LzmaIndex) LzmaIndexTotalSize() LzmaVli {
	return 0
}

/**
 * \brief       Get the total size of the file
 *
 * When no lzma_indexes have been combined with lzma_index_cat() and there is
 * no Stream Padding, this function is identical to lzma_index_stream_size().
 * If multiple lzma_indexes have been combined, this includes also the headers
 * of each separate Stream and the possible Stream Padding fields.
 *
 * \param       i   Pointer to lzma_index structure
 *
 * \return      Total size of the .xz file in bytes
 */
// llgo:link (*LzmaIndex).LzmaIndexFileSize C.lzma_index_file_size
func (recv_ *LzmaIndex) LzmaIndexFileSize() LzmaVli {
	return 0
}

/**
 * \brief       Get the uncompressed size of the file
 *
 * \param       i   Pointer to lzma_index structure
 *
 * \return      Size in bytes of the uncompressed data in the file
 */
// llgo:link (*LzmaIndex).LzmaIndexUncompressedSize C.lzma_index_uncompressed_size
func (recv_ *LzmaIndex) LzmaIndexUncompressedSize() LzmaVli {
	return 0
}

/**
 * \brief       Initialize an iterator
 *
 * This function associates the iterator with the given lzma_index, and calls
 * lzma_index_iter_rewind() on the iterator.
 *
 * This function doesn't allocate any memory, thus there is no
 * lzma_index_iter_end(). The iterator is valid as long as the
 * associated lzma_index is valid, that is, until lzma_index_end() or
 * using it as source in lzma_index_cat(). Specifically, lzma_index doesn't
 * become invalid if new Blocks are added to it with lzma_index_append() or
 * if it is used as the destination in lzma_index_cat().
 *
 * It is safe to make copies of an initialized lzma_index_iter, for example,
 * to easily restart reading at some particular position.
 *
 * \param       iter    Pointer to a lzma_index_iter structure
 * \param       i       lzma_index to which the iterator will be associated
 */
// llgo:link (*LzmaIndexIter).LzmaIndexIterInit C.lzma_index_iter_init
func (recv_ *LzmaIndexIter) LzmaIndexIterInit(i *LzmaIndex) {
}

/**
 * \brief       Rewind the iterator
 *
 * Rewind the iterator so that next call to lzma_index_iter_next() will
 * return the first Block or Stream.
 *
 * \param       iter    Pointer to a lzma_index_iter structure
 */
// llgo:link (*LzmaIndexIter).LzmaIndexIterRewind C.lzma_index_iter_rewind
func (recv_ *LzmaIndexIter) LzmaIndexIterRewind() {
}

/**
 * \brief       Get the next Block or Stream
 *
 * \param       iter    Iterator initialized with lzma_index_iter_init()
 * \param       mode    Specify what kind of information the caller wants
 *                      to get. See lzma_index_iter_mode for details.
 *
 * \return      lzma_bool:
 *              - true if no Block or Stream matching the mode is found.
 *                *iter is not updated (failure).
 *              - false if the next Block or Stream matching the mode was
 *                found. *iter is updated (success).
 */
// llgo:link (*LzmaIndexIter).LzmaIndexIterNext C.lzma_index_iter_next
func (recv_ *LzmaIndexIter) LzmaIndexIterNext(mode LzmaIndexIterMode) LzmaBool {
	return 0
}

/**
 * \brief       Locate a Block
 *
 * If it is possible to seek in the .xz file, it is possible to parse
 * the Index field(s) and use lzma_index_iter_locate() to do random-access
 * reading with granularity of Block size.
 *
 * If the target is smaller than the uncompressed size of the Stream (can be
 * checked with lzma_index_uncompressed_size()):
 *  - Information about the Stream and Block containing the requested
 *    uncompressed offset is stored into *iter.
 *  - Internal state of the iterator is adjusted so that
 *    lzma_index_iter_next() can be used to read subsequent Blocks or Streams.
 *
 * If the target is greater than the uncompressed size of the Stream, *iter
 * is not modified.
 *
 * \param       iter    Iterator that was earlier initialized with
 *                      lzma_index_iter_init().
 * \param       target  Uncompressed target offset which the caller would
 *                      like to locate from the Stream
 *
 * \return      lzma_bool:
 *              - true if the target is greater than or equal to the
 *                uncompressed size of the Stream (failure)
 *              - false if the target is smaller than the uncompressed size
 *                of the Stream (success)
 */
// llgo:link (*LzmaIndexIter).LzmaIndexIterLocate C.lzma_index_iter_locate
func (recv_ *LzmaIndexIter) LzmaIndexIterLocate(target LzmaVli) LzmaBool {
	return 0
}

/**
 * \brief       Concatenate lzma_indexes
 *
 * Concatenating lzma_indexes is useful when doing random-access reading in
 * multi-Stream .xz file, or when combining multiple Streams into single
 * Stream.
 *
 * \param[out]  dest      lzma_index after which src is appended
 * \param       src       lzma_index to be appended after dest. If this
 *                        function succeeds, the memory allocated for src
 *                        is freed or moved to be part of dest, and all
 *                        iterators pointing to src will become invalid.
* \param       allocator  lzma_allocator for custom allocator functions.
 *                        Set to NULL to use malloc() and free().
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: lzma_indexes were concatenated successfully.
 *                src is now a dangling pointer.
 *              - LZMA_DATA_ERROR: *dest would grow too big.
 *              - LZMA_MEM_ERROR
 *              - LZMA_PROG_ERROR
*/
// llgo:link (*LzmaIndex).LzmaIndexCat C.lzma_index_cat
func (recv_ *LzmaIndex) LzmaIndexCat(src *LzmaIndex, allocator *LzmaAllocator) LzmaRet {
	return 0
}

/**
 * \brief       Duplicate lzma_index
 *
 * \param       i         Pointer to lzma_index structure to be duplicated
 * \param       allocator lzma_allocator for custom allocator functions.
 *                        Set to NULL to use malloc() and free().
 *
 * \return      A copy of the lzma_index, or NULL if memory allocation failed.
 */
// llgo:link (*LzmaIndex).LzmaIndexDup C.lzma_index_dup
func (recv_ *LzmaIndex) LzmaIndexDup(allocator *LzmaAllocator) *LzmaIndex {
	return nil
}

/**
 * \brief       Initialize .xz Index encoder
 *
 * \param       strm        Pointer to properly prepared lzma_stream
 * \param       i           Pointer to lzma_index which should be encoded.
 *
 * The valid `action' values for lzma_code() are LZMA_RUN and LZMA_FINISH.
 * It is enough to use only one of them (you can choose freely).
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Initialization succeeded, continue with lzma_code().
 *              - LZMA_MEM_ERROR
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaStream).LzmaIndexEncoder C.lzma_index_encoder
func (recv_ *LzmaStream) LzmaIndexEncoder(i *LzmaIndex) LzmaRet {
	return 0
}

/**
 * \brief       Initialize .xz Index decoder
 *
 * \param       strm        Pointer to properly prepared lzma_stream
 * \param[out]  i           The decoded Index will be made available via
 *                          this pointer. Initially this function will
 *                          set *i to NULL (the old value is ignored). If
 *                          decoding succeeds (lzma_code() returns
 *                          LZMA_STREAM_END), *i will be set to point
 *                          to a new lzma_index, which the application
 *                          has to later free with lzma_index_end().
 * \param       memlimit    How much memory the resulting lzma_index is
 *                          allowed to require. liblzma 5.2.3 and earlier
 *                          don't allow 0 here and return LZMA_PROG_ERROR;
 *                          later versions treat 0 as if 1 had been specified.
 *
 * Valid `action' arguments to lzma_code() are LZMA_RUN and LZMA_FINISH.
 * There is no need to use LZMA_FINISH, but it's allowed because it may
 * simplify certain types of applications.
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Initialization succeeded, continue with lzma_code().
 *              - LZMA_MEM_ERROR
 *              - LZMA_PROG_ERROR
 *
 * \note        liblzma 5.2.3 and older list also LZMA_MEMLIMIT_ERROR here
 *              but that error code has never been possible from this
 *              initialization function.
 */
// llgo:link (*LzmaStream).LzmaIndexDecoder C.lzma_index_decoder
func (recv_ *LzmaStream) LzmaIndexDecoder(i **LzmaIndex, memlimit uint64) LzmaRet {
	return 0
}

/**
 * \brief       Single-call .xz Index encoder
 *
 * \note        This function doesn't take allocator argument since all
 *              the internal data is allocated on stack.
 *
 * \param       i         lzma_index to be encoded
 * \param[out]  out       Beginning of the output buffer
 * \param[out]  out_pos   The next byte will be written to out[*out_pos].
 *                        *out_pos is updated only if encoding succeeds.
 * \param       out_size  Size of the out buffer; the first byte into
 *                        which no data is written to is out[out_size].
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: Encoding was successful.
 *              - LZMA_BUF_ERROR: Output buffer is too small. Use
 *                lzma_index_size() to find out how much output
 *                space is needed.
 *              - LZMA_PROG_ERROR
 *
 */
// llgo:link (*LzmaIndex).LzmaIndexBufferEncode C.lzma_index_buffer_encode
func (recv_ *LzmaIndex) LzmaIndexBufferEncode(out *uint8, out_pos *uintptr, out_size uintptr) LzmaRet {
	return 0
}

/**
* \brief       Single-call .xz Index decoder
*
* \param[out]  i           If decoding succeeds, *i will point to a new
*                          lzma_index, which the application has to
*                          later free with lzma_index_end(). If an error
*                          occurs, *i will be NULL. The old value of *i
*                          is always ignored and thus doesn't need to be
*                          initialized by the caller.
* \param[out]  memlimit    Pointer to how much memory the resulting
*                          lzma_index is allowed to require. The value
*                          pointed by this pointer is modified if and only
*                          if LZMA_MEMLIMIT_ERROR is returned.
 * \param      allocator   lzma_allocator for custom allocator functions.
*                          Set to NULL to use malloc() and free().
* \param       in          Beginning of the input buffer
* \param       in_pos      The next byte will be read from in[*in_pos].
*                          *in_pos is updated only if decoding succeeds.
* \param       in_size     Size of the input buffer; the first byte that
*                          won't be read is in[in_size].
*
* \return      Possible lzma_ret values:
*              - LZMA_OK: Decoding was successful.
*              - LZMA_MEM_ERROR
*              - LZMA_MEMLIMIT_ERROR: Memory usage limit was reached.
*                The minimum required memlimit value was stored to *memlimit.
*              - LZMA_DATA_ERROR
*              - LZMA_PROG_ERROR
*/
//go:linkname LzmaIndexBufferDecode C.lzma_index_buffer_decode
func LzmaIndexBufferDecode(i **LzmaIndex, memlimit *uint64, allocator *LzmaAllocator, in *uint8, in_pos *uintptr, in_size uintptr) LzmaRet

/**
 * \brief       Initialize a .xz file information decoder
 *
 * This decoder decodes the Stream Header, Stream Footer, Index, and
 * Stream Padding field(s) from the input .xz file and stores the resulting
 * combined index in *dest_index. This information can be used to get the
 * uncompressed file size with lzma_index_uncompressed_size(*dest_index) or,
 * for example, to implement random access reading by locating the Blocks
 * in the Streams.
 *
 * To get the required information from the .xz file, lzma_code() may ask
 * the application to seek in the input file by returning LZMA_SEEK_NEEDED
 * and having the target file position specified in lzma_stream.seek_pos.
 * The number of seeks required depends on the input file and how big buffers
 * the application provides. When possible, the decoder will seek backward
 * and forward in the given buffer to avoid useless seek requests. Thus, if
 * the application provides the whole file at once, no external seeking will
 * be required (that is, lzma_code() won't return LZMA_SEEK_NEEDED).
 *
 * The value in lzma_stream.total_in can be used to estimate how much data
 * liblzma had to read to get the file information. However, due to seeking
 * and the way total_in is updated, the value of total_in will be somewhat
 * inaccurate (a little too big). Thus, total_in is a good estimate but don't
 * expect to see the same exact value for the same file if you change the
 * input buffer size or switch to a different liblzma version.
 *
 * Valid `action' arguments to lzma_code() are LZMA_RUN and LZMA_FINISH.
 * You only need to use LZMA_RUN; LZMA_FINISH is only supported because it
 * might be convenient for some applications. If you use LZMA_FINISH and if
 * lzma_code() asks the application to seek, remember to reset `action' back
 * to LZMA_RUN unless you hit the end of the file again.
 *
 * Possible return values from lzma_code():
 *   - LZMA_OK: All OK so far, more input needed
 *   - LZMA_SEEK_NEEDED: Provide more input starting from the absolute
 *     file position strm->seek_pos
 *   - LZMA_STREAM_END: Decoding was successful, *dest_index has been set
 *   - LZMA_FORMAT_ERROR: The input file is not in the .xz format (the
 *     expected magic bytes were not found from the beginning of the file)
 *   - LZMA_OPTIONS_ERROR: File looks valid but contains headers that aren't
 *     supported by this version of liblzma
 *   - LZMA_DATA_ERROR: File is corrupt
 *   - LZMA_BUF_ERROR
 *   - LZMA_MEM_ERROR
 *   - LZMA_MEMLIMIT_ERROR
 *   - LZMA_PROG_ERROR
 *
 * \param       strm        Pointer to a properly prepared lzma_stream
 * \param[out]  dest_index  Pointer to a pointer where the decoder will put
 *                          the decoded lzma_index. The old value
 *                          of *dest_index is ignored (not freed).
 * \param       memlimit    How much memory the resulting lzma_index is
 *                          allowed to require. Use UINT64_MAX to
 *                          effectively disable the limiter.
 * \param       file_size   Size of the input .xz file
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_MEM_ERROR
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaStream).LzmaFileInfoDecoder C.lzma_file_info_decoder
func (recv_ *LzmaStream) LzmaFileInfoDecoder(dest_index **LzmaIndex, memlimit uint64, file_size uint64) LzmaRet {
	return 0
}

type LzmaIndexHashS struct {
	Unused [8]uint8
}
type LzmaIndexHash LzmaIndexHashS

/**
 * \brief       Allocate and initialize a new lzma_index_hash structure
 *
 * If index_hash is NULL, this function allocates and initializes a new
 * lzma_index_hash structure and returns a pointer to it. If allocation
 * fails, NULL is returned.
 *
 * If index_hash is non-NULL, this function reinitializes the lzma_index_hash
 * structure and returns the same pointer. In this case, return value cannot
 * be NULL or a different pointer than the index_hash that was given as
 * an argument.
 *
 * \param       index_hash  Pointer to a lzma_index_hash structure or NULL.
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 *
 * \return      Initialized lzma_index_hash structure on success or
 *              NULL on failure.
 */
// llgo:link (*LzmaIndexHash).LzmaIndexHashInit C.lzma_index_hash_init
func (recv_ *LzmaIndexHash) LzmaIndexHashInit(allocator *LzmaAllocator) *LzmaIndexHash {
	return nil
}

/**
 * \brief       Deallocate lzma_index_hash structure
 *
 * \param       index_hash  Pointer to a lzma_index_hash structure to free.
 * \param       allocator   lzma_allocator for custom allocator functions.
 *                          Set to NULL to use malloc() and free().
 */
// llgo:link (*LzmaIndexHash).LzmaIndexHashEnd C.lzma_index_hash_end
func (recv_ *LzmaIndexHash) LzmaIndexHashEnd(allocator *LzmaAllocator) {
}

/**
 * \brief       Add a new Record to an Index hash
 *
 * \param       index_hash        Pointer to a lzma_index_hash structure
 * \param       unpadded_size     Unpadded Size of a Block
 * \param       uncompressed_size Uncompressed Size of a Block
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK
 *              - LZMA_DATA_ERROR: Compressed or uncompressed size of the
 *                Stream or size of the Index field would grow too big.
 *              - LZMA_PROG_ERROR: Invalid arguments or this function is being
 *                used when lzma_index_hash_decode() has already been used.
 */
// llgo:link (*LzmaIndexHash).LzmaIndexHashAppend C.lzma_index_hash_append
func (recv_ *LzmaIndexHash) LzmaIndexHashAppend(unpadded_size LzmaVli, uncompressed_size LzmaVli) LzmaRet {
	return 0
}

/**
 * \brief       Decode and validate the Index field
 *
 * After telling the sizes of all Blocks with lzma_index_hash_append(),
 * the actual Index field is decoded with this function. Specifically,
 * once decoding of the Index field has been started, no more Records
 * can be added using lzma_index_hash_append().
 *
 * This function doesn't use lzma_stream structure to pass the input data.
 * Instead, the input buffer is specified using three arguments. This is
 * because it matches better the internal APIs of liblzma.
 *
 * \param       index_hash      Pointer to a lzma_index_hash structure
 * \param       in              Pointer to the beginning of the input buffer
 * \param[out]  in_pos          in[*in_pos] is the next byte to process
 * \param       in_size         in[in_size] is the first byte not to process
 *
 * \return      Possible lzma_ret values:
 *              - LZMA_OK: So far good, but more input is needed.
 *              - LZMA_STREAM_END: Index decoded successfully and it matches
 *                the Records given with lzma_index_hash_append().
 *              - LZMA_DATA_ERROR: Index is corrupt or doesn't match the
 *                information given with lzma_index_hash_append().
 *              - LZMA_BUF_ERROR: Cannot progress because *in_pos >= in_size.
 *              - LZMA_PROG_ERROR
 */
// llgo:link (*LzmaIndexHash).LzmaIndexHashDecode C.lzma_index_hash_decode
func (recv_ *LzmaIndexHash) LzmaIndexHashDecode(in *uint8, in_pos *uintptr, in_size uintptr) LzmaRet {
	return 0
}

/**
 * \brief       Get the size of the Index field as bytes
 *
 * This is needed to verify the Backward Size field in the Stream Footer.
 *
 * \param       index_hash      Pointer to a lzma_index_hash structure
 *
 * \return      Size of the Index field in bytes.
 */
// llgo:link (*LzmaIndexHash).LzmaIndexHashSize C.lzma_index_hash_size
func (recv_ *LzmaIndexHash) LzmaIndexHashSize() LzmaVli {
	return 0
}

/**
 * \brief       Get the total amount of physical memory (RAM) in bytes
 *
 * This function may be useful when determining a reasonable memory
 * usage limit for decompressing or how much memory it is OK to use
 * for compressing.
 *
 * \return      On success, the total amount of physical memory in bytes
 *              is returned. If the amount of RAM cannot be determined,
 *              zero is returned. This can happen if an error occurs
 *              or if there is no code in liblzma to detect the amount
 *              of RAM on the specific operating system.
 */
//go:linkname LzmaPhysmem C.lzma_physmem
func LzmaPhysmem() uint64

/**
 * \brief       Get the number of processor cores or threads
 *
 * This function may be useful when determining how many threads to use.
 * If the hardware supports more than one thread per CPU core, the number
 * of hardware threads is returned if that information is available.
 *
 * \return      On success, the number of available CPU threads or cores is
 *              returned. If this information isn't available or an error
 *              occurs, zero is returned.
 */
//go:linkname LzmaCputhreads C.lzma_cputhreads
func LzmaCputhreads() uint32
